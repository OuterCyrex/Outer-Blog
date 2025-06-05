<template>
  <v-footer padless color="black">
    <v-col class="text-center white--text">
      <div style="font-size:20px" class="mb-2">网站已运行: {{this.timeString}}</div>
      <div>网站访问量: {{this.view}}</div>
      <a href="https://beian.miit.gov.cn/" target="_blank" style="color:white;margin-right:20px">蜀ICP备2024096882号</a>
      <div style="display: inline-block" class="ml-6rve">Copyright © - Outer Blog {{new Date().getFullYear()}}</div>
    </v-col>
  </v-footer>
</template>

<script>

export default {
  data(){
    return{
      RunTime:0,
      timeString:'00 天 00 时 00 分 00 秒',
      view:0,
    }
  },
  created() {
    this.timeClock()
    this.getView()
  },
  methods:{
    getRunTime(){
      let origin = new Date('2024-08-16 00:49:20').getTime()
      this.RunTime = new Date().getTime() - origin;
      let second = parseInt(this.RunTime / 1000);
      let minute = parseInt(second / 60);
      let hour = parseInt(minute / 60);
      let day = parseInt(hour / 24);
      this.timeString = `${day} 天 ` +  (hour%24).toString().padStart(2,'0') + " 时 " +  (minute%60).toString().padStart(2,'0') +" 分 " + (second%60).toString().padStart(2,'0')+ " 秒";
    },
    timeClock(){
      setInterval(this.getRunTime,1000)
    },
    async getView(){
      const {data:res}= await this.$http.get('article/' + 1)
      res.data.view ++
      this.view = res.data.view
      await this.$http.put('article/view/' + 1,res.data)
    },
  }
}
</script>
