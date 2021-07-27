import axios from 'axios';

process.env.NODE_TLS_REJECT_UNAUTHORIZED = '0';

export default axios.create({
  baseURL: process.env.API_URL || 'https://localhost:3000',
  headers: {
    'Content-Type': 'application/json; charset=utf-8',
  },
});
