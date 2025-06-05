import Vue from "vue";
import axios from "axios";

axios.defaults.baseURL = 'http://124.220.68.177:8080/api/v1';
Vue.prototype.$http = axios;
