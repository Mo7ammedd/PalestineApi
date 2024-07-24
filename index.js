const express = require('express');
const path = require('path');
const app = express();
const faviconPath = 'favicon.ico';
const cors = require('cors');
const casualtiesUrl = 'https://data.techforpalestine.org/api/v2/casualties_daily.min.json';
const killedUrl = 'https://data.techforpalestine.org/api/v2/killed-in-gaza.min.json';
const westbankUrl = 'https://data.techforpalestine.org/api/v2/west_bank_daily.min.json';

let casualtiesCsv, killedCsv, westbankCsv;

app.use(cors());
app.options('*', cors());

app.get('/', (req, res) => {
  res.send({ message: 'from the river to the sea' });
});

app.get('/killed', async (req, res) => {
  if (!killedCsv) {
    const fetch = (await import('node-fetch')).default;
    const response = await fetch(killedUrl);
    killedCsv = await response.json();
  }
  res.json(killedCsv);
});

app.get('/casualties', async (req, res) => {
  if (!casualtiesCsv) {
    const fetch = (await import('node-fetch')).default;
    const response = await fetch(casualtiesUrl);
    casualtiesCsv = await response.json();
  }
  res.json(casualtiesCsv);
});

app.get('/westbank', async (req, res) => {
  if (!westbankCsv) {
    const fetch = (await import('node-fetch')).default;
    const response = await fetch(westbankUrl);
    westbankCsv = await response.json();
  }
  res.json(westbankCsv);
});

app.get('/favicon.ico', (req, res) => {
  res.sendFile(path.join(__dirname, faviconPath));
});

// Start server
app.listen(3000, () => console.log(`Server running on port 3000`));
