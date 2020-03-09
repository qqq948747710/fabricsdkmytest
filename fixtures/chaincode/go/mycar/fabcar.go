package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
	"strconv"
)

type MyCar struct {

}
type Car struct {
	Make string
	Model string
	Colour string
	Owner string
}

func (t *MyCar) Init(stub shim.ChaincodeStubInterface) pb.Response{
	return shim.Success(nil)
}


func(t *MyCar)Invoke(stub shim.ChaincodeStubInterface)pb.Response{
	function,args:=stub.GetFunctionAndParameters()
	if function=="queryCar"{
	}else if function=="initLedger"{
		return t.initLedget(stub)
	}else if function=="createCar"{
		return t.createCar(stub,args)
	}else if function=="queryAllCars"{
		return t.queryAllCars(stub)
	}else if function=="changeCarOwner"{
		return t.changeCarOwner(stub,args)
	}
	return shim.Error("not this function!!!!")
}

func(t *MyCar)initLedget(stub shim.ChaincodeStubInterface)pb.Response{
	cars := []Car{
		Car{Make: "Toyota", Model: "Prius", Colour: "blue", Owner: "Tomoko"},
		Car{Make: "Ford", Model: "Mustang", Colour: "red", Owner: "Brad"},
		Car{Make: "Hyundai", Model: "Tucson", Colour: "green", Owner: "Jin Soo"},
		Car{Make: "Volkswagen", Model: "Passat", Colour: "yellow", Owner: "Max"},
		Car{Make: "Tesla", Model: "S", Colour: "black", Owner: "Adriana"},
		Car{Make: "Peugeot", Model: "205", Colour: "purple", Owner: "Michel"},
		Car{Make: "Chery", Model: "S22L", Colour: "white", Owner: "Aarav"},
		Car{Make: "Fiat", Model: "Punto", Colour: "violet", Owner: "Pari"},
		Car{Make: "Tata", Model: "Nano", Colour: "indigo", Owner: "Valeria"},
		Car{Make: "Holden", Model: "Barina", Colour: "brown", Owner: "Shotaro"},
	}
	i:=0
	for i<len(cars){
		fmt.Println("i is ",i)
		carAsBytes,_:=json.Marshal(cars[i])
		stub.PutState("CAR"+strconv.Itoa(i),carAsBytes)
		fmt.Println("Added",cars[i])
		i=i+1
	}
	return shim.Success(nil)
}

func(t *MyCar)createCar(stub shim.ChaincodeStubInterface,args []string)pb.Response{
	if len(args)!=5{
		return shim.Error("the args number is not.must be Carkey,Make,model,Colour,Owner")
	}
	var car =Car{Make:args[1],Model:args[2],Colour:args[3],Owner:args[4]}
	carAsBytes,_:=json.Marshal(car)
	stub.PutState(args[0],carAsBytes)
	return shim.Success(nil)
}

func(t *MyCar)queryAllCars(stub shim.ChaincodeStubInterface)pb.Response{
	startKey:="CAR0"
	endKey:="CAR999"
	resultsIterator,err:=stub.GetStateByRange(startKey,endKey)
	if err !=nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()
	var buffer bytes.Buffer
	buffer.WriteString("[")
	bArrayMemerAlredadyWritten:=false
	for resultsIterator.HasNext(){
		queryResponse,err:=resultsIterator.Next()
		if err!=nil{
			return shim.Error(err.Error())
		}
		if bArrayMemerAlredadyWritten == true{
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"Key\":")
		buffer.WriteString("\"")
		buffer.WriteString(queryResponse.Key)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Record\":")
		buffer.WriteString(string(queryResponse.Value))
		buffer.WriteString("}")
		bArrayMemerAlredadyWritten=true
	}
	buffer.WriteString("]")
	fmt.Printf("queryAllCar:\n%s\n",buffer.String())
	return shim.Success(buffer.Bytes())
}
func (t *MyCar)changeCarOwner(stub shim.ChaincodeStubInterface,args []string)pb.Response{
	if len(args)!=2{
		return shim.Error("args is not,must be 2")
	}
	carAsBytes,_:=stub.GetState(args[0])
	car:=Car{}
	json.Unmarshal(carAsBytes,&car)
	car.Owner=args[1]
	carAsBytes,_=json.Marshal(car)
	stub.PutState(args[0],carAsBytes)
	return shim.Success(nil)
}
func main(){
	err:=shim.Start(new(MyCar))
	if err!=nil{
		fmt.Printf("Error creating new Smart Contract: %s", err)
	}
}