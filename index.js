const express = require('express');
const axios = require('axios');
const path = require('path');

const app = express();
const faviconPath = 'favicon.ico';

const casualtiesUrl = 'https://data.techforpalestine.org/api/v2/casualties_daily.min.json';
const killedUrl = 'https://data.techforpalestine.org/api/v2/killed-in-gaza.min.json';
const westbankUrl = 'https://data.techforpalestine.org/api/v2/west_bank_daily.min.json';

let casualtiesCsv, killedCsv, westbankCsv;

app.use((req, res, next) => {
  res.setHeader('Access-Control-Allow-Origin', '*'); 
  next();
});

app.get('/', (req, res) => {
  res.send({ message: 'from the river to the see' });
});

app.get('/killed', async (req, res) => {
  if (!killedCsv) {
    const response = await axios.get(killedUrl);
    killedCsv = response.data;
  }
  res.json(killedCsv);
});

app.get('/casualties', async (req, res) => {
  if (!casualtiesCsv) {
    const response = await axios.get(casualtiesUrl);
    casualtiesCsv = response.data;
  }
  res.json(casualtiesCsv);
});

app.get('/westbank', async (req, res) => {
  if (!westbankCsv) {
    const response = await axios.get(westbankUrl);
    westbankCsv = response.data;
  }
  res.json(westbankCsv);
});

app.get('/favicon.ico', (req, res) => {
  res.sendFile(path.join(__dirname, faviconPath));
});

// Start server
const port = process.env.PORT || 3000;
app.listen(port, () => console.log(`Server running on port ${port}`));