package main 
import (
    "fmt"
    "io/ioutil"
    "net/http"	
    "encoding/json"
    "time"
    "strconv"
    "bufio"
    "os"
	"github.com/lvanneo/llog/llogger"

  
)


func GetInformation(server_id *int , areaid *int,server_name *string,page *string,query_order *string,kindid *string)(interface{}){
  
	resp, err := http.Get("https://recommd.xyq.cbg.163.com/cgi-bin/recommend.py?act=recommd_by_role&server_id="+strconv.Itoa(*server_id)+"&areaid="+strconv.Itoa(*areaid)+"&server_name="+*server_name+"&page="+*page+"&query_order="+*query_order+"&kindid="+*kindid)
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


func Jisuan( Qu string, Quid    float64 , Fwid   float64 ){
 
    server_id:= int(Quid)
    areaid:=int(Fwid)
    server_name:=Qu
    page,query_order,kindid :="1","unit_price+ASC","23"
 
  
   
    resp := GetInformation(&server_id,&areaid,&server_name,&page,&query_order,&kindid )
    result := resp.(map[string]interface{})
  
        arr := make([]interface{},0)
         for v,i:=range result{        
        if(v == "equips"){
           child :=i.([]interface{})
        
           for d:=range child{
            var dilbert Employee
         
             sun := child[d].(map[string]interface {})
            
             for a,e:= range sun{
                
                if(a=="selling_time"){
                    dilbert.selling_time = e.(string) 
                }else if(a=="desc"){
                    dilbert.desc =e.(string) 
                }else if(a =="server_name"){
                    dilbert.server_name = e.(string) 
                }else if(a=="area_name"){
                    dilbert.area_name =e.(string) 
                }else if(a=="price"){
                    dilbert.price = e.(string) 
                }
              
             }


             
             llogger.Info("----------------------正在请求","服务器",dilbert.area_name,dilbert.server_name,dilbert.selling_time,dilbert.price,dilbert.desc)
          
             arr= AppendByte(arr,dilbert)
           }
        }
        
        
    }

   
   

}
type Config struct {
    City []interface{}
}

type City struct {
    Fw      string
    Fwid    float64
    Qu    string
    Quid     float64
 
}



func testRead() interface{} {
    data, err := ioutil.ReadFile("./config.json")
    if err != nil { 
		return err.Error()
     }
     newList:=Config{}
     err=json.Unmarshal(data,&newList)
     if err!=nil{
        fmt.Println(err)
    }
 
     return newList.City
}
func AppendByte(slice []interface{}, data ...interface{}) []interface{} {
    m := len(slice)
    n := m + len(data)
    if n > cap(slice) { // if necessary, reallocate
        // allocate double what's needed, for future growth.
        newSlice := make([]interface{}, (n+1)*2)
        copy(newSlice, slice)
        slice = newSlice
    }
    slice = slice[0:n]
    copy(slice[m:n], data)
   
    return slice
}

func bianli(){
    resp :=testRead()
 
    result := resp.([]interface{})
   
    limiter := time.Tick(time.Millisecond * 1000)
    
  
     for i:=range result{     
        var dilbert City
        child :=result[i].(map[string]interface{})       
        for c,d:=range child{
         
                if(c=="Fw"){
                 
                    dilbert.Fw=d.(string)
                    }else if(c=="Fwid"){
                    
                        dilbert.Fwid =d.(float64) 
                    }else if(c =="Qu"){
                
                        dilbert.Qu = d.(string) 
                    }else if(c=="Quid"){
                    
                        dilbert.Quid =d.(float64) 
                    }
            
           }
         
           <-limiter
           Jisuan(dilbert.Qu,dilbert.Quid,dilbert.Fwid)
       
          
       
     }
       
     
}

func main(){

    
      bianli()
}