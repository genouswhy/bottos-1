package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"io/ioutil"
	"bytes"
	"errors"
	"io"
	"golang.org/x/net/context"
	"github.com/bottos-project/bottos/contract/abi"
	chain "github.com/bottos-project/bottos/api"
)

type BcliPushTrxInfo struct {
	sender string
	contract string
	method string
	ParamMap map[string]interface{}
}

func send_httpreq (get_or_post string, ReqUrl string, ReqMsg io.Reader) ([]byte, error) {
    var err error	
    //fmt.Println("send_httpreq 11: get_or_post:", get_or_post, ", ReqUrl:",ReqUrl, ", ReqMsg:", ReqMsg)
    
    client := &http.Client{}
    req, _ := http.NewRequest(get_or_post, ReqUrl, ReqMsg)
    //fmt.Println("send_httpreq 22: get_or_post:", get_or_post, ", ReqUrl:",ReqUrl, ", ReqMsg:", ReqMsg)
    //req.Header.Set("Connection", "keep-alive")
    resp, err := client.Do(req)
    
    if err != nil || resp == nil {
	fmt.Println("Error: get_or_post:", get_or_post, ", resp:", resp, ",err: ", err)
	return nil, errors.New("Error: send http failed")
    }	   
	
    if resp.StatusCode != 200 {
	fmt.Println("Error: get_or_post:", get_or_post, ", ReqUrl:",ReqUrl, ",resp: ", resp)
	return nil, errors.New("Error: resp.StatusCode is not 200")
    }

    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
		return nil, err
    }

    if resp.StatusCode != 200 {
		return nil, errors.New(string(body))
    }

    fmt.Println("body:", string(body))
    
    return body, nil
}

func (cli *CLI) BcliGetTransaction (http_method string, http_url string, trxhash string) {
	if http_method != "restful" && http_method != "grpc" {
		return
	}
	
	if http_method == "grpc" {
		gettrx := &chain.GetTransactionRequest{trxhash}
		newAccountRsp, err := cli.client.GetTransaction(context.TODO(), gettrx)
		if err != nil || newAccountRsp == nil {
			fmt.Println(err)
			return
		}

		if newAccountRsp.Errcode != 0 {
			fmt.Printf("Transfer error:\n")
			fmt.Printf("    %v\n", newAccountRsp.Msg)
			return
		}

		fmt.Printf("GetTransaction Succeed\n")

		b, _ := json.Marshal(newAccountRsp.Result)
		cli.jsonPrint(b)

		fmt.Printf("Trx: %v\n", newAccountRsp.Result)

	} else {
		gettrx := &chain.GetTransactionRequest{trxhash}
		req, _ := json.Marshal(gettrx)
		req_new := bytes.NewBuffer([]byte(req))
		fmt.Println("req_new:", req_new)
		httpRspBody, err := send_httpreq("POST", http_url, req_new)
		if err != nil || httpRspBody == nil {
			fmt.Println("Error. httpRspBody: ", httpRspBody, ", err: ", err)
			return
		}
		
		fmt.Println("Done. Get transaction result body is: ", httpRspBody, ".")

		var trx chain.Transaction
		err = json.Unmarshal(httpRspBody, &trx)
		if err != nil {
		    fmt.Println("Error! Unmarshal to trx failed: ", err, "| body is: ", httpRspBody, ".")
		    return
		}
		
		fmt.Println("Done. Get transaction result trx is: ", trx, ".")
	}
}

func (cli *CLI) BcliPushTransaction (http_method string, http_url string, pushtrxinfo *BcliPushTrxInfo) {
	if http_method != "restful" && http_method != "grpc" {
		return
	}

	Abi, abierr := getAbibyContractName(pushtrxinfo.contract)
        if abierr != nil {
           return
        }
	
	chainInfo, err := cli.getChainInfo()
	if err != nil {
		fmt.Println("QueryChainInfo error: ", err)
		return
	}
	
	mapstruct := make(map[string]interface{})
	
	for key, value := range(pushtrxinfo.ParamMap) {
	
        	abi.Setmapval(mapstruct, key, value)
	}

	param, _ := abi.MarshalAbiEx(mapstruct, &Abi, pushtrxinfo.contract, pushtrxinfo.method)

	trx := &chain.Transaction{
		Version:     1,
		CursorNum:   chainInfo.HeadBlockNum,
		CursorLabel: chainInfo.CursorLabel,
		Lifetime:    chainInfo.HeadBlockTime + 100,
		Sender:      pushtrxinfo.sender,
		Contract:    pushtrxinfo.contract,
		Method:      pushtrxinfo.method,
		Param:       BytesToHex(param),
		SigAlg:      1,
	}
	
	sign, err := cli.signTrx(trx, param)
	if err != nil {
		return
	}
	
	trx.Signature = sign
	var newAccountRsp *chain.SendTransactionResponse

	if http_method == "grpc" {
		newAccountRsp, err = cli.client.SendTransaction(context.TODO(), trx)
		if err != nil || newAccountRsp == nil {
			fmt.Println(err)
			return
		}
		if newAccountRsp.Errcode != 0 {
			fmt.Printf("Transfer error:\n")
			fmt.Printf("    %s\n", newAccountRsp)
			return
		}
	} else {
		req, _ := json.Marshal(trx)
    		req_new := bytes.NewBuffer([]byte(req))
		fmt.Println("req_new:", req_new)
		httpRspBody, err := send_httpreq("POST", http_url, req_new)
		if err != nil || httpRspBody == nil {
			fmt.Println("BcliPushTransaction Error:", err, ", httpRspBody: ", httpRspBody)
			return
		}
		var respbody chain.SendTransactionResponse
		json.Unmarshal(httpRspBody, &respbody)
		newAccountRsp = &respbody
		//fmt.Println("Done. Push transaction result body is: ", respbody, ".")
	}
	

	fmt.Printf("Transfer Succeed\n")
	fmt.Printf("Trx: \n")

	printTrx := Transaction{
		Version:     trx.Version,
		CursorNum:   trx.CursorNum,
		CursorLabel: trx.CursorLabel,
		Lifetime:    trx.Lifetime,
		Sender:      trx.Sender,
		Contract:    trx.Contract,
		Method:      trx.Method,
		Param:       BytesToHex(param),
		ParamBin:    trx.Param,
		SigAlg:      trx.SigAlg,
		Signature:   trx.Signature,
	}

	b, _ := json.Marshal(printTrx)
	cli.jsonPrint(b)
	fmt.Printf("TrxHash: %v\n", newAccountRsp.Result.TrxHash)
}

