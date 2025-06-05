<template>
  <v-card class="pb-10">
    <v-col class="pl-5 pr-5">
    <v-img src="https://img.oulu.me/019%20Malibu%20Beach.png" max-height="300px" class="rounded">
      <v-card-title class="d-flex justify-center">
          <v-card-title class="d-flex flex-row justify-center mb-5" style="user-select: none; font-size: calc(20px + 2vw); padding-top:10%; font-family:'Arial Rounded MT Bold',sans-serif;color:white">
            OuterCyrex Blog
          </v-card-title>
        <v-card-text class="d-flex flex-row justify-center" style="color:white;user-select: none; font-size: calc(20px + 0.5vw)">欢迎访问我的博客</v-card-text>
      </v-card-title>
    </v-img>
      <v-row class="pt-5 pr-3" rows="12">
        <v-col cols="5" v-if="this.isDateVisible">
          <v-card class="pa-0">
            <v-col class="pa-0">
              <v-date-picker width="100%" color="blue"></v-date-picker>
            </v-col>
          </v-card>
        </v-col>
        <v-col :cols="this.DateWidth" align-self="center" :class="this.DateSidebar">
          <v-row rows="12" class="mb-2">
            <v-col cols="12" class="pa-0">
              <v-card>
                <v-card-title>现在是北京时间：</v-card-title>
                <v-card-text><v-icon class="mr-2">{{'mdi-clock'}}</v-icon>{{this.timeString}}</v-card-text>
              </v-card>
            </v-col>
          </v-row>
          <v-row rows="8">
            <v-col cols="12" class="pa-0">
              <v-card @click="$router.push(`/detail/${LastArticleInfo.ID}`)" :max-height="this.lastArtMaxHeight">
                <v-col cols="12">
                  <v-card-text color="blue" :class="this.lastArticleMargin">
                    <div style="font-size:calc(20px + 1vw)">最新博客文章</div>
                  </v-card-text>
                  <v-card-title class="mb-2">
                    <v-chip color="pink" label class="white--text mr-2">{{this.LastArticleInfo.Category.name}}</v-chip>
                    {{this.LastArticleInfo.title}}
                  </v-card-title>
                  <v-card-subtitle>{{this.LastArticleInfo.desc}}</v-card-subtitle>
                  <v-divider></v-divider>
                  <v-card-text v-if="!isDateVisible || showTag">
                    <v-icon class="mr-2">{{'mdi-tag'}}</v-icon>
                    <span class="mr-10" style="color:deepskyblue;font-weight: bold;font-size: 16px">{{this.LastArticleInfo.Tag.name}}</span>
                    <v-icon class="mr-2 mb-0">{{'mdi-calendar-month'}}</v-icon>
                    <span>上传时间：{{this.LastArticleInfo.CreatedAt.slice(0,10) + ' ' + this.LastArticleInfo.CreatedAt.slice(11,19)}}</span>
                  </v-card-text>
                  <v-card-text></v-card-text>
                </v-col>
              </v-card>
            </v-col>
          </v-row>
        </v-col>
      </v-row>
      <v-divider></v-divider>
      <v-row rows="12" class="mt-2">
        <v-col cols="12">

          <v-card >
          <v-card class="pa-0" color="#f4f4f4" outlined>
            <div class="d-flex justify-center" style="background-color:#C5E1A5; color:white;font-size:calc(20px + 1vw);padding:10px">公共留言板</div>
            <CommentWeb :id="1"></CommentWeb>
          </v-card>

          </v-card>
        </v-col>
      </v-row>
    </v-col>
  </v-card>
</template>
<script>

import CommentWeb from "@/components/CommentWeb.vue";

export default {
  components: {CommentWeb},
  data(){
    return{
      LastArticleInfo:{},
      timeString:'00:00:00',
      viewportWidth:window.innerWidth,
      isDateVisible:false,
      DateWidth:'7',
      DateSidebar:'ml-3 pr-6',
      lastArticleMargin:'pa-4',
      articleInfo:{},
      lastArtMaxHeight:'261px',
      showTag:true
    }
  },
  created(){
    this.showDate()
    this.getArticleList()
    setInterval(this.updateClock,1000)
  },
  mounted() {
    window.addEventListener('resize',this.showDate)
  },
  methods:{
    async getArticleList(){
      const {data:res} = await this.$http.get("articles",{
        params:{
          name:'',
          pagesize:1,
          pagenum:1
        }
      })
      this.LastArticleInfo = res.data[0];
    },
    updateClock(){
    const now = new Date();
    let hours = now.getHours();
    let minutes = now.getMinutes();
    let seconds = now.getSeconds();
    hours = hours.toString().padStart(2, '0');
    minutes = minutes.toString().padStart(2, '0');
    seconds = seconds.toString().padStart(2, '0');
    this.timeString = `${hours}:${minutes}:${seconds}`;
    },
    showDate(){
      this.viewportWidth = window.innerWidth
      this.isDateVisible = this.viewportWidth > 900;
      this.DateWidth = this.isDateVisible ? '7' : '12'
      this.DateSidebar = this.isDateVisible ?'ma-0':'ml-3 pr-6'
      this.lastArticleMargin = this.viewportWidth > 1903 ?'pa-4':'pa-2'
      this.lastArtMaxHeight = this.isDateVisible ? '261px' : '500px'
      this.showTag = this.viewportWidth > 1904
    },
  }
}
</script>
