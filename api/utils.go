package api

import "net/rpc"

func CallRPC(url string, serviceMethod string, args interface{}) (*rpc.Client, *interface{}, error) {
	var response interface{}

	client, err := rpc.DialHTTP("tcp", url)
	if err != nil {
		return nil, nil, err
	}

	err = client.Call(serviceMethod, args, &response)
	if err != nil {
		return nil, nil, err
	}

	return client, &response, nil
}
