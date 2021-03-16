import axios from 'axios';
import getConfig from 'next/config';

const { publicRuntimeConfig } = getConfig();
const { API_URL } = publicRuntimeConfig;

const http = axios.create({
  baseURL: API_URL,
});

export default http;
