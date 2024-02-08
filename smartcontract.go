package chaincode

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// SmartContract provides functions for managing an Asset
type SmartContract struct {
	contractapi.Contract
}

// Asset describes basic details of what makes up a simple asset
// Insert struct field in alphabetic order => to achieve determinism across languages
// golang keeps the order when marshal to json but doesn't order automatically
type Asset struct {
	Address                    string  `json:"Address"`
	Brand                      string  `json:"Brand"`
	ChasisNo                   string  `json:"ChasisNo"`
	Color                      string  `json:"Color"`
	CylinderDisplacement       float32 `json:"CylinderDisplacement"`
	DocumentSerialNo           string  `json:"DocumentSerialNo"`
	FirstRegDate               string  `json:"FirstRegDate"`
	FuelType                   string  `json:"FuelType"`
	IDTaxNo                    string  `json:"IDTaxNo"`
	MaxLoadedWeight            float32 `json:"MaxLoadedWeight"`
	ModelYear                  int     `json:"ModelYear"`
	MotorNo                    string  `json:"MotorNo"`
	MotorPower                 float32 `json:"MotorPower"`
	Name                       string  `json:"Name"`
	NetWeight                  float32 `json:"NetWeight"`
	NotaryName                 string  `json:"NotaryName"`
	NotarySaleDate             string  `json:"NotarySaleDate"`
	NotarySaleNo               string  `json:"NotarySaleNo"`
	NumberOfSeats              int     `json:"NumberOfSeats"`
	NumberOfStandingPassengers int     `json:"NumberOfStandingPassengers"`
	OtherInfo                  string  `json:"OtherInfo"`
	PWRate                     float32 `json:"PWRate"`
	Plate                      string  `json:"Plate"`
	ProvinceCity               string  `json:"ProvinceCity"`
	PurposeOfUsage             string  `json:"PurposeOfUsage"`
	RegDate                    string  `json:"RegDate"`
	RegOrderNo                 string  `json:"RegOrderNo"`
	RightsInTheVehicle         string  `json:"RightsInTheVehicle"`
	SurnameTitle               string  `json:"SurnameTitle"`
	TradeName                  string  `json:"TradeName"`
	TrailerMaxLoadedWeight     float32 `json:"TrailerMaxLoadedWeight"`
	TypeApprovalNo             string  `json:"TypeApprovalNo"`
	VehicleClass               string  `json:"VehicleClass"`
	VehicleType                string  `json:"VehicleType"`
	WagonWeight                float32 `json:"WagonWeight"`
}

// InitLedger adds a base set of assets to the ledger
func (s *SmartContract) InitLedger(ctx contractapi.TransactionContextInterface) error {
	assets := []Asset{
		{
			Address:                    "123 Elm St, Springfield",
			Brand:                      "Toyota",
			ChasisNo:                   "YT1234ABCDEF",
			Color:                      "Red",
			CylinderDisplacement:       1.8,
			DocumentSerialNo:           "DSN123456789",
			FirstRegDate:               "2015-04-23",
			FuelType:                   "Petrol",
			IDTaxNo:                    "1234567890",
			MaxLoadedWeight:            1500,
			ModelYear:                  2015,
			MotorNo:                    "MT123456789",
			MotorPower:                 132,
			Name:                       "John Doe",
			NetWeight:                  1200,
			NotaryName:                 "Jane Smith",
			NotarySaleDate:             "2020-06-15",
			NotarySaleNo:               "NSN123456789",
			NumberOfSeats:              5,
			NumberOfStandingPassengers: 0,
			OtherInfo:                  "Well maintained, one owner",
			PWRate:                     95,
			Plate:                      "ABC123",
			ProvinceCity:               "Springfield",
			PurposeOfUsage:             "Personal",
			RegDate:                    "2015-05-01",
			RegOrderNo:                 "RON123456789",
			RightsInTheVehicle:         "None",
			SurnameTitle:               "Doe",
			TradeName:                  "Doe's Cars",
			TrailerMaxLoadedWeight:     500,
			TypeApprovalNo:             "TAN123456789",
			VehicleClass:               "Passenger",
			VehicleType:                "Sedan",
			WagonWeight:                1300,
		},
		{
			Address:                    "456 Maple Ave, Rivertown",
			Brand:                      "Honda",
			ChasisNo:                   "HN5678GHIJKL",
			Color:                      "Blue",
			CylinderDisplacement:       2.0,
			DocumentSerialNo:           "DSN987654321",
			FirstRegDate:               "2018-08-10",
			FuelType:                   "Diesel",
			IDTaxNo:                    "0987654321",
			MaxLoadedWeight:            1600,
			ModelYear:                  2018,
			MotorNo:                    "MT987654321",
			MotorPower:                 140,
			Name:                       "Alice Johnson",
			NetWeight:                  1300,
			NotaryName:                 "Bob Marley",
			NotarySaleDate:             "2021-07-20",
			NotarySaleNo:               "NSN987654321",
			NumberOfSeats:              4,
			NumberOfStandingPassengers: 0,
			OtherInfo:                  "Lease return, low mileage",
			PWRate:                     100,
			Plate:                      "XYZ789",
			ProvinceCity:               "Rivertown",
			PurposeOfUsage:             "Commercial",
			RegDate:                    "2018-09-05",
			RegOrderNo:                 "RON987654321",
			RightsInTheVehicle:         "Lease",
			SurnameTitle:               "Johnson",
			TradeName:                  "AJ Motors",
			TrailerMaxLoadedWeight:     600,
			TypeApprovalNo:             "TAN987654321",
			VehicleClass:               "Utility",
			VehicleType:                "SUV",
			WagonWeight:                1400,
		},
		{
			Address:                    "789 Birch Rd, Lakeside",
			Brand:                      "Ford",
			ChasisNo:                   "FD1234567890",
			Color:                      "White",
			CylinderDisplacement:       2.3,
			DocumentSerialNo:           "DSN555666777",
			FirstRegDate:               "2017-03-15",
			FuelType:                   "Electric",
			IDTaxNo:                    "5678901234",
			MaxLoadedWeight:            1700,
			ModelYear:                  2017,
			MotorNo:                    "MT555666777",
			MotorPower:                 150,
			Name:                       "Bruce Wayne",
			NetWeight:                  1400,
			NotaryName:                 "Clark Kent",
			NotarySaleDate:             "2022-05-25",
			NotarySaleNo:               "NSN555666777",
			NumberOfSeats:              5,
			NumberOfStandingPassengers: 0,
			OtherInfo:                  "Imported, premium package",
			PWRate:                     110,
			Plate:                      "DEF456",
			ProvinceCity:               "Lakeside",
			PurposeOfUsage:             "Personal",
			RegDate:                    "2017-04-10",
			RegOrderNo:                 "RON555666777",
			RightsInTheVehicle:         "None",
			SurnameTitle:               "Wayne",
			TradeName:                  "Wayne Enterprises",
			TrailerMaxLoadedWeight:     700,
			TypeApprovalNo:             "TAN555666777",
			VehicleClass:               "Passenger",
			VehicleType:                "Coupe",
			WagonWeight:                1450,
		},
		{
			Address:                    "101 Pine St, Hilltown",
			Brand:                      "Tesla",
			ChasisNo:                   "TS9876543210",
			Color:                      "Black",
			CylinderDisplacement:       0.0,
			DocumentSerialNo:           "DSN112233445",
			FirstRegDate:               "2020-01-20",
			FuelType:                   "Electric",
			IDTaxNo:                    "1122334455",
			MaxLoadedWeight:            1800,
			ModelYear:                  2020,
			MotorNo:                    "MT112233445",
			MotorPower:                 200,
			Name:                       "Tony Stark",
			NetWeight:                  1500,
			NotaryName:                 "Pepper Potts",
			NotarySaleDate:             "2023-04-30",
			NotarySaleNo:               "NSN112233445",
			NumberOfSeats:              4,
			NumberOfStandingPassengers: 0,
			OtherInfo:                  "Autopilot, full options",
			PWRate:                     120,
			Plate:                      "GHI789",
			ProvinceCity:               "Hilltown",
			PurposeOfUsage:             "Personal",
			RegDate:                    "2020-02-15",
			RegOrderNo:                 "RON112233445",
			RightsInTheVehicle:         "None",
			SurnameTitle:               "Stark",
			TradeName:                  "Stark Industries",
			TrailerMaxLoadedWeight:     800,
			TypeApprovalNo:             "TAN112233445",
			VehicleClass:               "Passenger",
			VehicleType:                "Sedan",
			WagonWeight:                1550,
		},
		{
			Address:                    "202 Oak Lane, Mountainview",
			Brand:                      "Volkswagen",
			ChasisNo:                   "VW1234567890A",
			Color:                      "Silver",
			CylinderDisplacement:       1.6,
			DocumentSerialNo:           "DSN556677889",
			FirstRegDate:               "2019-07-29",
			FuelType:                   "Hybrid",
			IDTaxNo:                    "2233445566",
			MaxLoadedWeight:            1450,
			ModelYear:                  2019,
			MotorNo:                    "MT556677889",
			MotorPower:                 130,
			Name:                       "Emma Watson",
			NetWeight:                  1250,
			NotaryName:                 "Daniel Radcliffe",
			NotarySaleDate:             "2024-01-10",
			NotarySaleNo:               "NSN556677889",
			NumberOfSeats:              5,
			NumberOfStandingPassengers: 0,
			OtherInfo:                  "Eco-friendly, low emissions",
			PWRate:                     90,
			Plate:                      "JKL012",
			ProvinceCity:               "Mountainview",
			PurposeOfUsage:             "Personal",
			RegDate:                    "2019-08-15",
			RegOrderNo:                 "RON556677889",
			RightsInTheVehicle:         "None",
			SurnameTitle:               "Watson",
			TradeName:                  "Watson's Eco Rides",
			TrailerMaxLoadedWeight:     550,
			TypeApprovalNo:             "TAN556677889",
			VehicleClass:               "Passenger",
			VehicleType:                "Hatchback",
			WagonWeight:                1350,
		},
	}

	for _, asset := range assets {
		assetJSON, err := json.Marshal(asset)
		if err != nil {
			return err
		}

		err = ctx.GetStub().PutState(asset.ChasisNo, assetJSON)
		if err != nil {
			return fmt.Errorf("failed to put to world state. %v", err)
		}
	}

	return nil
}

// CreateAsset issues a new asset to the world state with given details.
func (s *SmartContract) CreateAsset(
	ctx contractapi.TransactionContextInterface,
	address string,
	brand string,
	chasis_no string,
	color string,
	cylinder_displacement float32,
	document_serial_no string,
	first_reg_date string,
	fuel_type string,
	id_tax_no string,
	max_loaded_weight float32,
	model_year int,
	motor_no string,
	motor_power float32,
	name string,
	net_weight float32,
	notary_name string,
	notary_sale_date string,
	notary_sale_no string,
	number_of_seats int,
	number_of_standing_passengers int,
	other_info string,
	pw_rate float32,
	plate string,
	province_city string,
	purpose_of_usage string,
	reg_date string,
	reg_order_no string,
	rights_in_the_vehicle string,
	surname_title string,
	trade_name string,
	trailer_max_loaded_weight float32,
	type_approval_no string,
	vehicle_class string,
	vehicle_type string,
	wagon_weight float32,

) error {
	exists, err := s.AssetExists(ctx, chasis_no)
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("the asset %s already exists", chasis_no)
	}

	asset := Asset{
		Address:                    address,
		Brand:                      brand,
		ChasisNo:                   chasis_no,
		Color:                      color,
		CylinderDisplacement:       cylinder_displacement,
		DocumentSerialNo:           document_serial_no,
		FirstRegDate:               first_reg_date,
		FuelType:                   fuel_type,
		IDTaxNo:                    id_tax_no,
		MaxLoadedWeight:            max_loaded_weight,
		ModelYear:                  model_year,
		MotorNo:                    motor_no,
		MotorPower:                 motor_power,
		Name:                       name,
		NetWeight:                  net_weight,
		NotaryName:                 notary_name,
		NotarySaleDate:             notary_sale_date,
		NotarySaleNo:               notary_sale_no,
		NumberOfSeats:              number_of_seats,
		NumberOfStandingPassengers: number_of_standing_passengers,
		OtherInfo:                  other_info,
		PWRate:                     pw_rate,
		Plate:                      plate,
		ProvinceCity:               province_city,
		PurposeOfUsage:             purpose_of_usage,
		RegDate:                    reg_date,
		RegOrderNo:                 reg_order_no,
		RightsInTheVehicle:         rights_in_the_vehicle,
		SurnameTitle:               surname_title,
		TradeName:                  trade_name,
		TrailerMaxLoadedWeight:     trailer_max_loaded_weight,
		TypeApprovalNo:             type_approval_no,
		VehicleClass:               vehicle_class,
		VehicleType:                vehicle_type,
		WagonWeight:                wagon_weight,
	}
	assetJSON, err := json.Marshal(asset)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(chasis_no, assetJSON)
}

// ReadAsset returns the asset stored in the world state with given id.
func (s *SmartContract) ReadAsset(
	ctx contractapi.TransactionContextInterface, chasis_no string) (*Asset, error) {
	assetJSON, err := ctx.GetStub().GetState(chasis_no)
	if err != nil {
		return nil, fmt.Errorf("failed to read from world state: %v", err)
	}
	if assetJSON == nil {
		return nil, fmt.Errorf("the asset %s does not exist", chasis_no)
	}

	var asset Asset
	err = json.Unmarshal(assetJSON, &asset)
	if err != nil {
		return nil, err
	}

	return &asset, nil
}

// UpdateAsset updates an existing asset in the world state with provided parameters.
func (s *SmartContract) UpdateAsset(
	ctx contractapi.TransactionContextInterface,
	address string,
	brand string,
	chasis_no string,
	color string,
	cylinder_displacement float32,
	document_serial_no string,
	first_reg_date string,
	fuel_type string,
	id_tax_no string,
	max_loaded_weight float32,
	model_year int,
	motor_no string,
	motor_power float32,
	name string,
	net_weight float32,
	notary_name string,
	notary_sale_date string,
	notary_sale_no string,
	number_of_seats int,
	number_of_standing_passengers int,
	other_info string,
	pw_rate float32,
	plate string,
	province_city string,
	purpose_of_usage string,
	reg_date string,
	reg_order_no string,
	rights_in_the_vehicle string,
	surname_title string,
	trade_name string,
	trailer_max_loaded_weight float32,
	type_approval_no string,
	vehicle_class string,
	vehicle_type string,
	wagon_weight float32,

) error {
	exists, err := s.AssetExists(ctx, chasis_no)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("the asset %s does not exist", chasis_no)
	}

	// overwriting original asset with new asset
	asset := Asset{
		Address:                    address,
		Brand:                      brand,
		ChasisNo:                   chasis_no,
		Color:                      color,
		CylinderDisplacement:       cylinder_displacement,
		DocumentSerialNo:           document_serial_no,
		FirstRegDate:               first_reg_date,
		FuelType:                   fuel_type,
		IDTaxNo:                    id_tax_no,
		MaxLoadedWeight:            max_loaded_weight,
		ModelYear:                  model_year,
		MotorNo:                    motor_no,
		MotorPower:                 motor_power,
		Name:                       name,
		NetWeight:                  net_weight,
		NotaryName:                 notary_name,
		NotarySaleDate:             notary_sale_date,
		NotarySaleNo:               notary_sale_no,
		NumberOfSeats:              number_of_seats,
		NumberOfStandingPassengers: number_of_standing_passengers,
		OtherInfo:                  other_info,
		PWRate:                     pw_rate,
		Plate:                      plate,
		ProvinceCity:               province_city,
		PurposeOfUsage:             purpose_of_usage,
		RegDate:                    reg_date,
		RegOrderNo:                 reg_order_no,
		RightsInTheVehicle:         rights_in_the_vehicle,
		SurnameTitle:               surname_title,
		TradeName:                  trade_name,
		TrailerMaxLoadedWeight:     trailer_max_loaded_weight,
		TypeApprovalNo:             type_approval_no,
		VehicleClass:               vehicle_class,
		VehicleType:                vehicle_type,
		WagonWeight:                wagon_weight,
	}
	assetJSON, err := json.Marshal(asset)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(chasis_no, assetJSON)
}

// DeleteAsset deletes an given asset from the world state.
func (s *SmartContract) DeleteAsset(ctx contractapi.TransactionContextInterface, chasis_no string) error {
	exists, err := s.AssetExists(ctx, chasis_no)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("the asset %s does not exist", chasis_no)
	}

	return ctx.GetStub().DelState(chasis_no)
}

// AssetExists returns true when asset with given ID exists in world state
func (s *SmartContract) AssetExists(ctx contractapi.TransactionContextInterface, chasis_no string) (bool, error) {
	assetJSON, err := ctx.GetStub().GetState(chasis_no)
	if err != nil {
		return false, fmt.Errorf("failed to read from world state: %v", err)
	}

	return assetJSON != nil, nil
}

// GetAllAssets returns all assets found in world state
func (s *SmartContract) GetAllAssets(ctx contractapi.TransactionContextInterface) ([]*Asset, error) {
	// range query with empty string for startKey and endKey does an
	// open-ended query of all assets in the chaincode namespace.
	resultsIterator, err := ctx.GetStub().GetStateByRange("", "")
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var assets []*Asset
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		var asset Asset
		err = json.Unmarshal(queryResponse.Value, &asset)
		if err != nil {
			return nil, err
		}
		assets = append(assets, &asset)
	}

	return assets, nil
}
