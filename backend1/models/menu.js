const mongoose=require("mongoose")

const menuSchema=new mongoose.Schema({
    name: { type: String, required: true},
    description: {type: String, required: true},
    price: {type: String},
    
})

module.exports=mongoose.model("menu", menuSchema)