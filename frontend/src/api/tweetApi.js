import axios from 'axios';

const BASE_URL = 'http://localhost:8080/api';

const tweetApi = axios.create({
  baseURL: BASE_URL,
  headers: {
    'Content-Type': 'application/json',
  },
  withCredentials: true,
});

export const getTweetList = async () => {
  try {
    const response = await tweetApi.get(`/tweetList`);
    return response.data;
  } catch (error) {
    throw error.response.data;
  }
};