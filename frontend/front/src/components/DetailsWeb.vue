<template>
  <div>
  <v-card class="pb-10">
    <div class="d-flex justify-center text-h4 font-weight-bold pt-3 pb-4 pl-3">{{articleInfo.title}}</div>
    <div class="d-flex justify-center pb-1" style="font-size:18px">作者: {{this.author}}</div>
    <div class="d-flex justify-center pb-4" style="font-size:15px">点击量: {{this.articleInfo.view}}</div>
    <v-divider class="pa-3 ma-3"></v-divider>
    <v-alert class="ma-4" elevation="1" color="indigo" dark border="left" outlined>{{articleInfo.desc}}</v-alert>
    <v-card-text>
      <div class="ArticleHTML" v-html="articleInfo.html"></div>
    </v-card-text>
    <div class="d-flex justify-center" style="font-size:20px;user-select: none">----- end -----</div>
    <v-row justify="space-between" no-gutters class="ml-3 mr-3">
      <v-col cols="4" class="ma-4">
        <v-card style="user-select: none" @click=this.pushAfter :disabled="!this.articleAfter.html" class="pa-5 pt-2" :min-height="pageHeight">
          <div style="color:black;" class="mb-2 align-end">← 下一篇<v-icon>{{'mdi-right'}}</v-icon></div>
          <div style="color:black;font-weight: bold;font-size: 18px" v-if="this.viewportWidth>600">{{this.articleAfter.title}}</div>
        </v-card>
      </v-col>
      <v-col cols="4" class="ma-4">
        <v-card class="pa-5 pt-2" @click=this.pushBefore :disabled="!this.articleBefore.html" :min-height="pageHeight" >
          <div style="color:black;" class="mb-2">上一篇 →<v-icon>{{'mdi-right'}}</v-icon></div>
          <div style="color:black;font-weight: bold;font-size: 18px" v-if="this.viewportWidth>600">{{this.articleBefore.title}}</div>
        </v-card>
      </v-col>
    </v-row>
    <v-divider></v-divider>
    <v-card :class="commentWidth">
      <v-card class="pa-0" color="#f4f4f4" outlined>
        <div class="d-flex justify-center" style="background-color: #00E5FF; color:white;font-size:calc(20px + 1vw);padding:10px">文章评论区</div>
        <CommentWeb :id="this.id"></CommentWeb>
      </v-card>
    </v-card>
  </v-card>
  </div>
</template>

<script>
import CommentWeb from "@/components/CommentWeb.vue";
window.onload = function() {
  window.scrollTo(0,500);
};
export default {
  components: {CommentWeb},
  props:['id'],
  data(){
    return{
      articleInfo:{},
      author:'',
      viewportWidth:1000,
      commentWidth:'ma-10',
      articleBefore:{title:"未查询到上一篇文章"},
      articleAfter:{title:"未查询到下一篇文章"},
      pageHeight:'120px'
    }
  },
  created(){
    this.getArticleInfo()
    this.changeSize()
  },
  mounted() {
    window.addEventListener('resize',this.changeSize)
  },
  methods:{
    pushAfter(){
      this.$router.push(`/detail/${this.articleAfter.ID}`)
      window.location.reload();
      window.scrollTo(0,0)
    },
    pushBefore(){
      this.$router.push(`/detail/${this.articleBefore.ID}`)
      window.location.reload();
      window.scrollTo(0,0);
    },
    changeSize(){
      this.viewportWidth = window.innerWidth
      this.commentWidth = this.viewportWidth > 900 ? 'ma-10':''
      this.pageHeight = this.viewportWidth > 600 ? '120px':'40px'
    },
    async getArticleInfo(){
      const {data:res}= await this.$http.get('article/' + this.id)
      res.data.view ++
      this.articleInfo = res.data
      await this.$http.put('article/view/' + this.id,res.data)
      this.getUsername()
      this.articleBefore.ID = this.articleInfo.ID - 1
      this.articleAfter.ID = this.articleInfo.ID + 1
      const {data:resBefore}= await this.$http.get('article/' + this.articleBefore.ID)
      if(resBefore.status !== 3001){
        this.articleBefore = resBefore.data
      }else{
        this.articleBefore.title = "未查询到上一篇文章"
      }
      const {data:resAfter}= await this.$http.get('article/' + this.articleAfter.ID)
      if(resAfter.status !== 3001){
        this.articleAfter = resAfter.data
      }else{
        this.articleAfter.title = "未查询到下一篇文章"
      }
    },
    async getUsername() {
      const {data:res} = await this.$http.get('user/' + this.articleInfo.uid)
      this.author = res.data.username
    },
  }
}
</script>

<style scoped>
.ArticleHTML{
  color:black;
}
.ArticleHTML >>> pre{
  font-size:20px;
  font-family:"Cascadia Code",sans-serif;
  margin:20px 10px;
  color:white;
  background-color:#333;
  padding:20px;
  width:90%;
  border-radius:6px;
  overflow-x: auto;
}
.ArticleHTML >>> .hljs-keyword{
  color:orange;
}
.ArticleHTML >>> .hljs-title{
  color:cornflowerblue;
}
.ArticleHTML >>> .hljs-type{
  color:palegoldenrod;
}
.ArticleHTML >>> .hljs-string{
  color:orangered;
}
.ArticleHTML >>> .hljs-literal{
  color:pink;
}
.ArticleHTML >>> .hljs-params{
  color:mediumpurple;
}
.ArticleHTML >>> .hljs-comment{
  color:greenyellow;
}
.ArticleHTML >>> .hljs-number{
  color:yellow;
}
.ArticleHTML >>> span{
  font-family:"Cascadia Code",sans-serif;
}
.ArticleHTML >>> h1{
  margin:20px 0;
}
.ArticleHTML >>> h2{
  margin:15px 0;
}
.ArticleHTML >>> h3{
  margin:10px 0;
}
.ArticleHTML >>> code{
  font-family: "Cascadia Code",sans-serif;
  font-size:12px;
}
.ArticleHTML >>> td{
  padding:10px;
  border: 1px solid black;
  border-collapse: collapse;
}
.ArticleHTML >>> table{
  margin:40px auto;
  padding:4px;
  border: 1px solid black;
  border-collapse: collapse;
}
.ArticleHTML >>> th{
  padding:10px;
  border: 1px solid black;
  border-collapse: collapse;
}
.ArticleHTML >>> img{
  width:auto;
  height: auto;
  max-width: 90%;
}
.ArticleHTML >>> hr{
  margin-top:20px;
  margin-bottom:20px;
}
.ArticleHTML >>> .hljs-section{
  color:deepskyblue;
}
.ArticleHTML >>> .hljs-attr{
  color:deepskyblue;
}
</style>
