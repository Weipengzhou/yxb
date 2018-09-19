package main 
import (
    "fmt"
    "io/ioutil"
    "net/http"	
    "encoding/json"

)


func GetInformation(server_id *string , areaid *string,server_name *string,page *string,query_order *string,kindid *string)(interface{}){
	
	resp, err := http.Get("https://recommd.xyq.cbg.163.com/cgi-bin/recommend.py?act=recommd_by_role&server_id="+*server_id+"&areaid="+*areaid+"&server_name="+*server_name+"&page="+*page+"&query_order="+*query_order+"&kindid="+*kindid)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil { 
		return err.Error()
     }
     result := map[string]interface{}{}
     json.Unmarshal(body, &result)
  

	return result
}
type Employee struct {  
    selling_time      string
    desc              string
    server_name       string
    area_name         string
    price             string
}

func Jisuan(){
    server_id,areaid,server_name,page,query_order,kindid :="554","58","兰亭序","1","unit_price+ASC","23"
    resp := GetInformation(&server_id,&areaid,&server_name,&page,&query_order,&kindid )
    result := resp.(map[string]interface{})
  
    for v,i:=range result{        
        if(v == "equips"){
           child :=i.([]interface{})
           for d:=range child{
             sun := child[d].(map[string]interface {})
             for a,e:= range sun{
                fmt.Println(a,e)
             }
           }
        }
          
    }
}
func main(){
   
      Jisuan()

}