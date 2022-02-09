import axios from 'axios'

const { REACT_APP_SERVER_URL: serverURL } = process.env

const axiosInstance = axios.create() 
axiosInstance.defaults.baseURL = serverURL
console.log(serverURL)

export default axiosInstance