<template>
  <v-main style="background-color: rgba(0, 0, 0, 0.5);height:100vh">
    <v-container style="height:100vh;margin-top:20vh">
    <v-row justify="center">
      <v-col cols="4" class="my-custom-col">
        <v-card v-if="!isLoggedIn&&!isSignVisible" style="border-radius:10px">
        <div style="padding: 60px 50px 30px 50px;">
          <v-form ref="formRef" v-model="valid" lazy-validation>
        <v-text-field label="用户名" type="username" v-model="userInfo.username" :rules="nameRules"></v-text-field>
        <v-text-field label="密码" v-on:keyup.enter="login" type="password" v-model="userInfo.password" :rules="pwdRules"></v-text-field>

          <div style="display:flex;justify-content: flex-end">

            <v-btn text color="blue" class="mr-10" @click="isSignVisible = true">注册</v-btn>
          <v-btn @click="clear" color="error" style="margin-right:20px;color:black">取消</v-btn>
          <v-btn style="background-color: deepskyblue;color:white" @click="login">登录</v-btn>
            </div>
              </v-form>
          </div>
        </v-card>
        <v-card v-if="!isLoggedIn&&isSignVisible" style="border-radius:10px">
          <div style="padding: 60px 50px 30px 50px;">
            <v-form ref="newFormRef" v-model="valid" lazy-validation>
              <v-text-field label="用户名" type="username" v-model="newUser.newUsername" :rules="nameRules"></v-text-field>
              <v-text-field label="密码" type="password" v-model="newUser.newPassword" :rules="pwdRules"></v-text-field>
              <v-text-field class="mb-4" label="确认密码" v-on:keyup.enter="login" type="password" v-model="newUser.checkPwd" :rules="checkpwdRules"></v-text-field>
              <span style="color:#e33f3f">请保存好自己的用户名和密码</span>
              <div style="display:flex;justify-content: flex-end" class="mt-5">
                <v-btn @click="clearSign" color="error" style="margin-right:20px;color:black">取消</v-btn>
                <v-btn style="background-color: deepskyblue;color:white" @click="handleSign" :loading="this.signTimeout !== null">注册</v-btn>
              </div>
            </v-form>
          </div>
        </v-card>
        <v-card v-if="isLoggedIn" style="border-radius:10px">
          <v-card-title class="ma-0">登录信息</v-card-title>
          <div style="padding: 0 50px;">
            <div style="font-size:20px;" class="mb-2">用户名：{{this.user.username}}</div>
            <span  style="font-size:20px;">权限：</span>
            <v-chip width="100%" :color="this.user.role === 1 ? 'pink':'blue'" label class="white--text">{{this.user.role === 1? '管理员' : '用户'}}</v-chip>
          </div>
          <v-btn class="ma-5" @click="quitLogin" color="error" style="margin-right:20px;color:black">退出登录</v-btn>
          <v-btn class="ml-2" @click="quitInfo" color="primary" style="margin-right:20px;color:black">确认</v-btn>
        </v-card>
      </v-col>
    </v-row>
    <v-snackbar
      v-model="snackbar.value"
      :timeout="2000"
      top
      style="color:white"
    >{{snackbar.text}}</v-snackbar>
  </v-container>
  </v-main>
</template>

<script>

export default {
  data(){
    return{
      valid: true,
      userInfo:{
        username:'',
        password:'',
      },
      newUser:{
        newUsername:'',
        newPassword:'',
        checkPwd:'',
      },
      snackbar: {
        value: false,
        text: '',
      },
      nameRules: [
        v => !!v || '请输入用户名',
        v => (v && v.length <= 12 && v.length>=4) || '用户名应为 4 至 12 个字符',
      ],
      pwdRules:[
        v => !!v || '请输入密码',
        v => (v && v.length <= 12 && v.length>=4) || '密码应为 6 至 20 个字符',
      ],
      checkpwdRules:[
        v => !!v || '请再次输入密码',
        v =>  (v && v === this.newUser.newPassword) || '两次密码不一致',
      ],
      loginValid:false,
      user:{},
      isLoggedIn:false,
      isSignVisible:false,
      signTimeout: null,
    }
  },
  created(){
    this.getUserInfo()
  },
  methods:{
    handleSign() {
      if (this.signTimeout) {
        clearTimeout(this.signTimeout);
      }
      this.signTimeout = setTimeout(() => {
        this.sign()
        this.signTimeout = null;
      }, 2000);
    },
    clear(){
      this.$refs.formRef.reset()
      this.$emit('update:data', this.loginValid);
    },
    clearSign(){
      this.$refs.newFormRef.reset()
      this.$emit('update:data', this.loginValid);
    },
    quitInfo(){
      this.$emit('update:data', this.loginValid);
    },
    quitLogin(){
      window.sessionStorage.setItem("token",'');
      this.isLoggedIn = false;
      this.$emit('update:isLoggedIn', this.isLoggedIn);
      this.$emit('update:data', this.loginValid);
    },
    async login(){
      try {
        const isValid = await this.$refs.formRef.validate();
        if (!isValid) return;

        const response = await this.$http.post("userlogin", this.userInfo);
        if (response.data.status !== 200) {
          this.snackbar.value = true
          this.snackbar.text = response.data.msg;
          return;
        }
        window.sessionStorage.setItem("token", response.data.token);
        this.isLoggedIn = true;
        this.$emit('update:isLoggedIn', this.isLoggedIn);
        this.$emit('update:data', this.loginValid);
      } catch (error) {
        this.snackbar.value = true
        this.snackbar.text = error;
      }
    },
    async sign(){
      try {
        const isValid = await this.$refs.newFormRef.validate();
        if (!isValid) return;
        let newUser = {
          username:this.newUser.newUsername,
          password:this.newUser.newPassword,
          role:2,
        }
        const response = await this.$http.post("user/add",newUser);
        if (response.data.status !== 200) {
          this.snackbar.value = true
          this.snackbar.text = response.data.message;
        }else{
          this.snackbar.value = true
          this.snackbar.text = "注册成功";
          this.isSignVisible = false;
        }
      } catch (error) {
        this.snackbar.value = true
        this.snackbar.text = error;
      }
    },
    async getUserInfo(){
      const token = window.sessionStorage.getItem("token");

      if(!token)return;

      const config = {
        headers: {
          'Authorization': `Bearer ${token}`
        }
      };

      const {data:res} = await this.$http.get("userinfo",config)

      if(res.status!==200){
        this.isLoggedIn = false;
        return
      }
      this.isLoggedIn = true;
      this.user = res.data;
    },
  }
}
</script>

<style scoped>
.my-custom-col {
  min-width: 350px;
  flex: 0 0 auto;
}
</style>

