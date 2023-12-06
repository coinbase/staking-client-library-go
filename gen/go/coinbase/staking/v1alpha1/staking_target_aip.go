// Code generated by protoc-gen-go-aip. DO NOT EDIT.
//
// versions:
// 	protoc-gen-go-aip development
// 	protoc (unknown)
// source: coinbase/staking/v1alpha1/staking_target.proto

package v1alpha1

import (
	fmt "fmt"
	resourcename "go.einride.tech/aip/resourcename"
	strings "strings"
)

type ValidatorResourceName struct {
	Protocol  string
	Network   string
	Validator string
}

func (n ProtocolResourceName) ValidatorResourceName(
	network string,
	validator string,
) ValidatorResourceName {
	return ValidatorResourceName{
		Protocol:  n.Protocol,
		Network:   network,
		Validator: validator,
	}
}

func (n NetworkResourceName) ValidatorResourceName(
	validator string,
) ValidatorResourceName {
	return ValidatorResourceName{
		Protocol:  n.Protocol,
		Network:   n.Network,
		Validator: validator,
	}
}

func (n ValidatorResourceName) Validate() error {
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
	if n.Validator == "" {
		return fmt.Errorf("validator: empty")
	}
	if strings.IndexByte(n.Validator, '/') != -1 {
		return fmt.Errorf("validator: contains illegal character '/'")
	}
	return nil
}

func (n ValidatorResourceName) ContainsWildcard() bool {
	return false || n.Protocol == "-" || n.Network == "-" || n.Validator == "-"
}

func (n ValidatorResourceName) String() string {
	return resourcename.Sprint(
		"protocols/{protocol}/networks/{network}/stakingTargets/{validator}",
		n.Protocol,
		n.Network,
		n.Validator,
	)
}

func (n ValidatorResourceName) MarshalString() (string, error) {
	if err := n.Validate(); err != nil {
		return "", err
	}
	return n.String(), nil
}

func (n *ValidatorResourceName) UnmarshalString(name string) error {
	return resourcename.Sscan(
		name,
		"protocols/{protocol}/networks/{network}/stakingTargets/{validator}",
		&n.Protocol,
		&n.Network,
		&n.Validator,
	)
}

func (n ValidatorResourceName) ProtocolResourceName() ProtocolResourceName {
	return ProtocolResourceName{
		Protocol: n.Protocol,
	}
}

func (n ValidatorResourceName) NetworkResourceName() NetworkResourceName {
	return NetworkResourceName{
		Protocol: n.Protocol,
		Network:  n.Network,
	}
}

type ContractResourceName struct {
	Protocol string
	Network  string
	Contract string
}

func (n ProtocolResourceName) ContractResourceName(
	network string,
	contract string,
) ContractResourceName {
	return ContractResourceName{
		Protocol: n.Protocol,
		Network:  network,
		Contract: contract,
	}
}

func (n NetworkResourceName) ContractResourceName(
	contract string,
) ContractResourceName {
	return ContractResourceName{
		Protocol: n.Protocol,
		Network:  n.Network,
		Contract: contract,
	}
}

func (n ContractResourceName) Validate() error {
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
	if n.Contract == "" {
		return fmt.Errorf("contract: empty")
	}
	if strings.IndexByte(n.Contract, '/') != -1 {
		return fmt.Errorf("contract: contains illegal character '/'")
	}
	return nil
}

func (n ContractResourceName) ContainsWildcard() bool {
	return false || n.Protocol == "-" || n.Network == "-" || n.Contract == "-"
}

func (n ContractResourceName) String() string {
	return resourcename.Sprint(
		"protocols/{protocol}/networks/{network}/stakingTargets/{contract}",
		n.Protocol,
		n.Network,
		n.Contract,
	)
}

func (n ContractResourceName) MarshalString() (string, error) {
	if err := n.Validate(); err != nil {
		return "", err
	}
	return n.String(), nil
}

func (n *ContractResourceName) UnmarshalString(name string) error {
	return resourcename.Sscan(
		name,
		"protocols/{protocol}/networks/{network}/stakingTargets/{contract}",
		&n.Protocol,
		&n.Network,
		&n.Contract,
	)
}

func (n ContractResourceName) ProtocolResourceName() ProtocolResourceName {
	return ProtocolResourceName{
		Protocol: n.Protocol,
	}
}

func (n ContractResourceName) NetworkResourceName() NetworkResourceName {
	return NetworkResourceName{
		Protocol: n.Protocol,
		Network:  n.Network,
	}
}
