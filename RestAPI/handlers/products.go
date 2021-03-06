package handlers
import (
	"log"
	"net/http"
  "encoding/json"
)
type Products struct {
l *log.Logger
}
func NewProducts(l*log.Logger) *Products{
	return &Products{l}
}
func (p *Products) ServeHttp(rw  http.ResponseWriter , r *http.Request) {
	if r.Method == http.MethodGet {
		p.getProducts(rw,r)
		return 
	}
	if r.Method ==http.MethodPost {
		p.addProduct(rw, r)
		return
	}
	rw.WriteHeader(http.StatusMethodNotAllowed)
}
	
func (p *Products) getProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET Products")

	lp :=data.GetProducts()

err := lp.ToJSON(rw)
if err! = nil {
	http.Error(rw, " unable to marshal json", http.StatusInternalServerError)
}

}
func (p *Products) addProduct(rw http.ResponseWriter, r*http.Request){
	p.l.Println("Handle POST Products")

	prod :=&data.Product{}
	err := prod.FromJSON(r.Body)

	if err !=nil{
		http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
	}
	data.addProduct(prod)
}