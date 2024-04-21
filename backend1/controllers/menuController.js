const menu=require('../models/menu')

async function createMenu(req, res) {
    try {
      const { name, description, price } = req.body;
  
      if (!name || !description) {
        return res.status(400).json({ message: 'Name and description are required' });
      }
  
      const newMenu = new menu({ name, description, price });
      const savedMenu = await newMenu.save();
  
      res.status(201).json(savedMenu);
    } catch (err) {
      console.error(err);
      res.status(500).json({ message: 'Server error' });
    }
  }

async function getController(req,res){
    try{
      const dataAll = await menu.find()   
      if(dataAll.length==0){
          return res.status(404).json({
          message: "No data found",
        })
      }
      else {
        return res.status(201).json(dataAll);
      }
    }
    catch(err){
      console.error(err);
      return res.status(500).send(err);
    }
  }
  module.exports={
     getController,
     createMenu,
  };