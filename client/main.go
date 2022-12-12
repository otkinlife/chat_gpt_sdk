package main

import (
	"flag"
	"fmt"
	"os"
	"syscall"

	"github.com/otkinlife/chat_gpt_sdk"
	"golang.org/x/term"
)

var (
	oldTermState *term.State
)

const prompt = "ChatGPT> "

func main() {
	var token string
	flag.StringVar(&token, "t", "", "【chatGPT token】run code in chrome console with JSON.parse(document.getElementById('__NEXT_DATA__').text).props.pageProps.accessToken")
	flag.Parse()
	token = "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCIsImtpZCI6Ik1UaEVOVUpHTkVNMVFURTRNMEZCTWpkQ05UZzVNRFUxUlRVd1FVSkRNRU13UmtGRVFrRXpSZyJ9.eyJodHRwczovL2FwaS5vcGVuYWkuY29tL3Byb2ZpbGUiOnsiZW1haWwiOiJpYWJ5Y3h2MDU3MzlAMjFjbi5jb20iLCJlbWFpbF92ZXJpZmllZCI6dHJ1ZSwiZ2VvaXBfY291bnRyeSI6IlVTIn0sImh0dHBzOi8vYXBpLm9wZW5haS5jb20vYXV0aCI6eyJ1c2VyX2lkIjoidXNlci16NWNXVzVoRjJESHdUSXZpR3lpYjJCeFAifSwiaXNzIjoiaHR0cHM6Ly9hdXRoMC5vcGVuYWkuY29tLyIsInN1YiI6ImF1dGgwfDYzOTQxYjBkZmU2OGQ0YjQ5MWI3NWVmMCIsImF1ZCI6WyJodHRwczovL2FwaS5vcGVuYWkuY29tL3YxIiwiaHR0cHM6Ly9vcGVuYWkuYXV0aDAuY29tL3VzZXJpbmZvIl0sImlhdCI6MTY3MDgxNDAxMSwiZXhwIjoxNjcwODU3MjExLCJhenAiOiJUZEpJY2JlMTZXb1RIdE45NW55eXdoNUU0eU9vNkl0RyIsInNjb3BlIjoib3BlbmlkIGVtYWlsIHByb2ZpbGUgbW9kZWwucmVhZCBtb2RlbC5yZXF1ZXN0IG9yZ2FuaXphdGlvbi5yZWFkIG9mZmxpbmVfYWNjZXNzIn0.OL7GXYq2UdW5ts7rNo7X0vE9U14ZqDZ_ADCwCFeSOFiHNRmpkKFju3WUmFA7Z0d7iccJUmKa-pqUNL2Nt-8w6o2SYKGLDYbeC82-oSIHZvuu-3v8nVve8HUM3XWT9LZe1_gL46T_QgULfLCVdrsCVrxPIGT0JW5kT0RfHJGqr5buArCiXIS0Bri2DFt11ou9322c_Enw6rrXxsK_3TlJAe2t9q264hRWDYG6iQ21gtDHKtd6uPdc-eMKHuB7kG3C1oPZPap8L9MmdVqekxH7PNpC34_n_SAXM61Ax_H9FUOYYzsSZUaC5gvXHodlF61mwnSeqsYVNUa7j-cgFslo_A"
	if token == "" {
		panic("token is empty")
	}

	var err error
	oldTermState, err = term.MakeRaw(syscall.Stdin)
	if err != nil {
		fmt.Println(err)
		return
	}

	t := term.NewTerminal(os.Stdin, prompt)
	t.Write(t.Escape.Blue)
	t.Write([]byte("ChatGPT CLI <KcJia> - version 1.0\n"))
	if err != nil {
		fmt.Println(err)
		return
	}

	chat := chat_gpt_sdk.NewChatGPT(token, t)

	for {
		t.SetPrompt("Question> ")
		question, err := t.ReadLine()
		if err != nil {
			fmt.Println(err.Error())
			break
		}
		t.Write([]byte("Answer> "))
		err = chat.Talk(question)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
	}
	term.Restore(syscall.Stdin, oldTermState)
}
