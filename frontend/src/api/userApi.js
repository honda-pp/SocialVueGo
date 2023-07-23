import axios from 'axios';

const BASE_URL = 'http://localhost:8080/api';

const userApi = axios.create({
  baseURL: BASE_URL,
  headers: {
    'Content-Type': 'application/json',
  },
});

export const loginUser = async (userData) => {
  try {
    const response = await userApi.post(`/login`, userData);
    return response.data;
  } catch (error) {
    throw error.response.data;
  }
};