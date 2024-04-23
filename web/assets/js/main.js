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
                    const id = Number(e.target.id);
                    let count = this.order.get(id);
                    if(count == NaN || count == null || count == undefined)
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
                    window.sessionStorage.order = JSON.stringify(Array.from(this.order.entries()));
                }
        },
        created(){
                    fetch("http://localhost:8000/api/prices", {
                        method: "GET",
                    })
                    .then((response) => response.json())
                    .then((response) => this.products = new Map(Object.entries(response)))
                    .catch((err) => console.log(err));
                 
                    // let str = '{"1": { "id":1, "title":"mam", "description":"sadas", "price":1231}, "2": { "id":2, "title":"mam", "description":"sadas", "price":1231}}'
                    // this.products = new Map(Object.entries(JSON.parse(str)))


                    window.addEventListener("storage", function(e){
                        console.log(event);
                    })

                    console.log(this.products);
                if(window.sessionStorage.order == null){
                    this.order = new Map()   
                    return;
                }

                this.order = new Map(JSON.parse(window.sessionStorage.order));
                this.sumProducts();
        }
})