
<template>
  <div>
    <v-app-bar app color="indigo darken-2" >
      <v-container class="py-0 fill-height">
        <v-avatar @click="LoginCouldSee" size="40" class="mr-3" color="grey">
          <v-img v-if="Logged" src="../assets/defaultAvatar.png"></v-img>
        </v-avatar>
        <v-btn text color="white" @click="$router.push('/')">首页</v-btn>
        <v-btn text color="white" @click="$router.push('/articles')">文章列表</v-btn>
        <div v-if="this.isTagVisible">
        <v-btn v-for="item in categoryList" :key="item.id" text color="white" @click="$router.push(`/articles/${item.id}`)">{{item.name}}</v-btn>
        </div>
        <div v-if="!this.isTagVisible">
        <v-btn text color="white" @click="this.showList" class="ChooseTag">选择分类</v-btn>
          <v-list v-if="isListVisible" style="position: absolute; top:50px;" rounded color="indigo" class="TagList">
            <v-list-item style="color:white" v-for="item in categoryList" :key="item.id" @click="$router.push(`/articles/${item.id}`)">{{item.name}}</v-list-item>
          </v-list>
        </div>
      </v-container>
        <v-spacer></v-spacer>
        <v-responsive v-if="isSearchVisible" min-width="300" max-width="400" color="white">
          <v-text-field dense flat hide-details solo-inverted rounded dark
                        placeholder="根据标题查询"
                        v-model="queryParam.name"
                        v-on:keyup.enter="sendDataToParent"
                        append-icon="mdi-close"
                        @click:append="clearInput"
          ></v-text-field>
        </v-responsive>
    </v-app-bar>
  </div>
</template>

<script>
export default {
  props:['isLoggedIn'],
  data(){
    return{
      queryParam:{
        name:'',
        pagesize:5,
        pagenum:1,
      },
      categoryList:[],
      Logged : false,
      isSearchVisible:false,
      viewportWidth:window.innerWidth,
      isTagVisible:true,
      isListVisible:false,
    }
  },
  created() {
    this.GetCategoryList()
    this.getUserInfo()
    this.showSearch()
  },
  watch:{
    isLoggedIn(newVal){
      this.Logged = newVal
    }
  },
  mounted() {
    window.addEventListener('resize',this.showSearch);
    window.addEventListener('click',this.HideList);
  },
  methods:{
    async GetCategoryList(){
      const {data:res}= await this.$http.get('categories',{
        params:
            {
              name:'',
              pagesize:1000,
              pagenum:1,
            },
      });
      this.categoryList=res.data;
    },
    sendDataToParent() {
      this.$router.push('/articles')
      this.$emit('update:data', this.queryParam.name);
    },
    clearInput(){
      this.queryParam.name=''
      this.sendDataToParent()
    },
    LoginCouldSee(){
      this.$emit('update:Login', true);
    },
    showSearch(){
     this.viewportWidth = window.innerWidth;
     this.isSearchVisible = this.viewportWidth > 995;
     this.isTagVisible = this.viewportWidth > 1263;
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
        this.Logged = false;
        return
      }
      this.Logged = true;
    },
    showList(){
      this.isListVisible=true
    },
    HideList(event) {
      const targetElement = this.$el.querySelector('.TagList');
      const chooseTag = this.$el.querySelector('.ChooseTag')
      if(chooseTag!=null&&targetElement!=null){
        if (!targetElement.contains(event.target) &&!chooseTag.contains(event.target)) {
          this.isListVisible = false;
        }
      }
    },
  }
}
</script>
