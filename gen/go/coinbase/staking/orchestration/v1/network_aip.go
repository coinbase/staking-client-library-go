// Code generated by protoc-gen-go-aip. DO NOT EDIT.
//
// versions:
// 	protoc-gen-go-aip development
// 	protoc (unknown)
// source: coinbase/staking/orchestration/v1/network.proto

package v1

import (
	fmt "fmt"
	resourcename "go.einride.tech/aip/resourcename"
	strings "strings"
)

type NetworkResourceName struct {
	Protocol string
	Network  string
}

func (n ProtocolResourceName) NetworkResourceName(
	network string,
) NetworkResourceName {
	return NetworkResourceName{
		Protocol: n.Protocol,
		Network:  network,
	}
}

func (n NetworkResourceName) Validate() error {
	if n.Protocol == "" {
		return fmt.Errorf("protocol: empty")
	}
	if strings.IndexByte(n.Protocol, '/') != -1 {
		return fmt.Errorf("protocol: contains illegal character '/'")
	}
	if n.Network == "" {
		return fmt.Errorf("network: empty")
	}
	if strings.IndexByte(n.Network, '/') != -1 {
		return fmt.Errorf("network: contains illegal character '/'")
	}
	return nil
}

func (n NetworkResourceName) ContainsWildcard() bool {
	return false || n.Protocol == "-" || n.Network == "-"
}

func (n NetworkResourceName) String() string {
	return resourcename.Sprint(
		"protocols/{protocol}/networks/{network}",
		n.Protocol,
		n.Network,
	)
}

func (n NetworkResourceName) MarshalString() (string, error) {
	if err := n.Validate(); err != nil {
		return "", err
	}
	return n.String(), nil
}

func (n *NetworkResourceName) UnmarshalString(name string) error {
	return resourcename.Sscan(
		name,
		"protocols/{protocol}/networks/{network}",
		&n.Protocol,
		&n.Network,
	)
}

func (n NetworkResourceName) ProtocolResourceName() ProtocolResourceName {
	return ProtocolResourceName{
		Protocol: n.Protocol,
	}
}
