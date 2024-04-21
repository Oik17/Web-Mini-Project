const express = require('express');
const dotenv = require('dotenv');
const cors = require('cors');
const mongoose = require('mongoose');
const { createMenu,getController } = require('./controllers/menuController');

dotenv.config();
mongoose.connect(process.env.DBURI);

const app = express();
const port = process.env.PORT;

app.use((req, res, next) => {
  res.header('Access-Control-Allow-Origin', '*');
  res.header('Access-Control-Allow-Headers', 'Origin, X-Requested-With, Content-Type, Accept');
  next();
});



app.use(express.json());
app.use(cors());
app.get("/ping", (req, res) => {
  res.send("pong");
});

app.post("/create-food", createMenu) 
app.get("/get-menu", getController)
app.listen(port, () => {
  console.log(`Listening on port ${port}`);
});