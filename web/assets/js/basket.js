const app = new Vue({
                el:"#app",
                data:{
                    products:new Map(),
	                sum: 0,
                    order:new Map(),
	                productsOrder:[],
                    //model product:
                    // product: {
                    //     id:0,
                    //     image:"",
                    //     title:"",
                    //     description:"",
                    //     price:""
                    // }
                },
                methods:{
                    sumPrice: function(){
                        let sum = 0;
                        this.productsOrder.forEach(v => {
                            if(this.order.get(v.id) != undefined){
                                sum += v.price*this.order.get(v.id);
                            }
                        });
                        this.sum = sum;
                    },
                	addProduct: function(e){
                		const id = Number(e.target.id);
                		let count = this.order.get(id);
                		this.order.set(id, count+1);

                		this.saveLocalOrder();
                        this.sumPrice();
                        this.$forceUpdate();
                	},
                    removeItem:function(id){
                        document.getElementById(id+"-con").remove();
                    },
                    removeProduct: function(e){
                        const id = Number(e.target.id);
                        if(id != undefined && this.order.get(id) <= 1){
                            this.order.set(id, 0);
                            this.order.delete(id);
                            this.removeItem(id);
                        }else{
                            let count = this.order.get(id);
                            this.order.set(id, count-1);
                        }

                        this.sumPrice();
                    	this.saveLocalOrder();
                        this.$forceUpdate();
                    },
                    saveLocalOrder: function(){
                        window.sessionStorage.order = JSON.stringify(Array.from(this.order.entries()));
                    },
                    setOrder: function(){
                        //can use cache system
                        this.order.forEach((value, key) => {
                            this.productsOrder.push(this.products.get(key+""));
                            console.log(value + "" + key)
                        })
                        
                    },
                    postOrder: function(e){
                        window.sessionStorage.clear();
                        this.$forceUpdate();
                    }
                },
                mounted(){
                    fetch("http://localhost:8000/api/prices", {
                        method: "GET",
                    })
                    .then((response) => response.json())
                    .then((response) => this.products = response
                    )
                    
                    .catch((err) => console.log(err));

                    // let str = '{"1": { "id":1, "title":"mam", "description":"sadas", "price":1231, "op":12}, "2": { "id":2, "title":"mam", "description":"sadas", "price":1231}}'
                    // this.products = new Map(Object.entries(JSON.parse(str)))

                    console.log(this.products);
                    
                        //TEST CASE
                        // //
                        //     this.order = new Map()
                        //     this.order.set(1,12);
                        //     this.products.set(1, {
                        //         id:1,
                        //         title:"name",
                        //         description:"asdasd",
                        //         price:123
                        //     })
                        //     localStorage.order = JSON.stringify(Array.from(this.order.entries()))
                        // //


                    if(window.sessionStorage.order == null){
                    	this.order = new Map();
                    	return;
                    }

                    console.log(JSON.parse(window.sessionStorage.order))
                    this.setOrder();
                    this.sumPrice();
                }
              })