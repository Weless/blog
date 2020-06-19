<template>
  <div class="widget widget-recent-posts">
    <h3 class="widget-title">最新文章</h3>
    <ul>
      <li v-for="(article,index) in articleList" :key="'latest'+index">
        <a :href="'/article/'+article.ArticleId"> {{ article.Title}}</a>
      </li>
    </ul>
  </div>
</template>

<script>
import service from "@/utils/request"
export default {
    name:"Latest",
    data(){
      return {
        articleList:[]
      }
    },
    methods:{
      getLatestArticles(){
        service.get("/latest")
        .then(res=>{
          // console.log(res)
          this.articleList = res.data.data.articleList
        })
      }
    },
    created(){
      this.getLatestArticles()
    }
}
</script>