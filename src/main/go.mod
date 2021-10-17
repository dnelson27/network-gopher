module dnelson-infosec.com/network-gopher/main

go 1.13

replace dnelson-infosec.com/network-gopher/graph => ../graph

replace dnelson-infosec.com/network-gopher/traversal => ../traversal

replace dnelson-infosec.com/network-gopher/utils => ../utils

require (
	dnelson-infosec.com/network-gopher/graph v0.0.0-00010101000000-000000000000
	dnelson-infosec.com/network-gopher/traversal v0.0.0-00010101000000-000000000000
	dnelson-infosec.com/network-gopher/utils v0.0.0-00010101000000-000000000000 // indirect
)
