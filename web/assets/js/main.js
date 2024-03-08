const app = new Vue({
    el:"#app",
        data:{
            products:new Map(),
            order:new Map(),
            sum: 0,
            product: {
                id:0,
                image:"",
                title:"",
                description:"",
                price:""
                }
            },
    methods:{
        sumProducts: function(){
                    let sum = 0;
                    this.order.forEach( (v,k) => {
                        sum+=v;
                    });

                    this.sum = sum;
                },
        addProduct: function(e){
                    const id = e.target.id;
                    let count = this.order.get(e.target.id);
                    if(count == NaN || count== undefined)
                        count = 0;
                    this.order.set(id, count+1);

                    this.sumProducts();
                    this.saveLocalOrder();
                },
        removeProduct: function(e){
                    this.order.set(e.target.id, "");

                    this.sumProducts();
                    this.saveLocalOrder();
                },
        saveLocalOrder: function(){
                    localStorage.order = JSON.stringify(Array.from(this.order.entries()));
                }
        },
        created(){
                 
                       fetch('http://localhost:8000/api/prices', {
                        method: "GET",
                        headers: {
                          //
                        },
                      })
                        .then((response) => {
                          response.json().then((data) => {
                            this.products = new Map(Object.entries(data));
                          });
                        })
                        .catch((err) => {
                          console.error(err);
                        });
                console.log(this.products)
                if(localStorage.order == null){
                    this.order = new Map()   
                    return;
                }
                
                this.order = new Map(JSON.parse(localStorage.order));
                this.sumProducts();
        }
})