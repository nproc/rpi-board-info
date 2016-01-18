package gql

import (
	"github.com/graphql-go/graphql"
	"github.com/nproc/rpi-board-info/info"
)

func NewSchema(
	board info.BoardInformationProvider,
	cpu info.CPUInformationProvider,
	disk info.DiskInformationProvider,
	load info.LoadAverageInformationProvider,
	mem info.MemoryInformationProvider,
	net info.NetworkInformationProvider,
	tmp info.TemperatureInformationProvider,
) (graphql.Schema, error) {
	return graphql.NewSchema(
		graphql.SchemaConfig{
			Query: graphql.NewObject(
				graphql.ObjectConfig{
					Name: "Root",
					Fields: graphql.Fields{
						"board": &graphql.Field{
							Type: Board(),
							Resolve: func(p graphql.ResolveParams) (interface{}, error) {
								board, err := board.Board()
								if err != nil {
									return nil, err
								}
								return NewBoardData(board), nil
							},
						},
						"cpu": &graphql.Field{
							Type: CPU(),
							Resolve: func(p graphql.ResolveParams) (interface{}, error) {
								stat, err := cpu.CPU()
								if err != nil {
									return nil, err
								}
								return NewCPUData(stat), nil
							},
						},
						"disks": &graphql.Field{
							Type: graphql.NewList(Disk()),
							Resolve: func(p graphql.ResolveParams) (interface{}, error) {
								disks, err := disk.Disks()
								if err != nil {
									return nil, err
								}
								return NewDiskDataList(disks), nil
							},
						},
						"load": &graphql.Field{
							Type: LoadAverage(),
							Resolve: func(p graphql.ResolveParams) (interface{}, error) {
								stat, err := load.LoadAverage()
								if err != nil {
									return nil, err
								}
								return NewLoadAverageData(stat), nil
							},
						},
						"memory": &graphql.Field{
							Type: graphql.NewObject(
								graphql.ObjectConfig{
									Name: "Memory",
									Fields: graphql.Fields{
										"swap": &graphql.Field{
											Type: Swap(),
											Resolve: func(p graphql.ResolveParams) (interface{}, error) {
												swap, err := mem.SwapMemory()
												if err != nil {
													return nil, err
												}
												return NewSwapData(swap), nil
											},
										},
										"virtual": &graphql.Field{
											Type: Virtual(),
											Resolve: func(p graphql.ResolveParams) (interface{}, error) {
												virtual, err := mem.VirtualMemory()
												if err != nil {
													return nil, err
												}
												return NewVirtualData(virtual), nil
											},
										},
									},
								},
							),
							Resolve: func(p graphql.ResolveParams) (interface{}, error) {
								return NewMemoryData(), nil
							},
						},
						"network": &graphql.Field{
							Type: graphql.NewList(Network()),
							Resolve: func(p graphql.ResolveParams) (interface{}, error) {
								stats, err := net.Networks()
								if err != nil {
									return nil, err
								}
								return NewNetworkDataList(stats), nil
							},
						},
						"temperature": &graphql.Field{
							Type: graphql.NewObject(
								graphql.ObjectConfig{
									Name: "Temperature",
									Fields: graphql.Fields{
										"cpu": &graphql.Field{
											Type: CPUTemperature(),
											Resolve: func(p graphql.ResolveParams) (interface{}, error) {
												stat, err := tmp.CPUTemperature()
												if err != nil {
													return nil, err
												}
												return NewTemperatureData(stat), nil
											},
										},
									},
								},
							),
							Resolve: func(p graphql.ResolveParams) (interface{}, error) {
								return NewTemperatureRoot(), nil
							},
						},
					},
				},
			),
		},
	)
}
