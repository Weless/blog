<template>
  <div>
    <article class="post post-4" v-for="(article,index) in articleList" :key="'article-'+index">
      <header class="entry-header">
        <h1 class="entry-title">
          <a :href="'/article/' + article.ArticleId">{{article.Title}}</a>
        </h1>
        <div class="entry-meta">
          <span class="post-category">
            <a href="javascript:void(0);">{{ article.Name }}</a>
          </span>
          <span class="post-date">
            <a href="javascript:void(0);">
              <time class="entry-date" datetime="20121109" v-html="dealWithTime(article.CreateTime)"></time>
            </a>
          </span>
          <!-- <span class="post-author">
          <a href="#">追梦人物</a>
          </span>-->
          <span class="comments-link">
            <a href="javascript:void(0);">{{ article.CommentCount}} 评论</a>
          </span>
          <span class="views-count">
            <a href="javascript:void(0);">{{ article.ViewCount}} 阅读</a>
          </span>
        </div>
      </header>
      <div class="entry-content clearfix">
        <p>{{ article.Summary}}</p>
        <div class="read-more cl-effect-14">
          <a :href="'/article/' + article.ArticleId" class="more-link">
            继续阅读
            <span class="meta-nav">→</span>
          </a>
        </div>
      </div>
    </article>
    <pagination :paginationData="panigationData"></pagination>
  </div>
  
</template>

<script>
import Pagination from "./Pagination.vue"
import service from "@/utils/request"
export default {

  name: "Article",
  components:{
    Pagination
  },
  data() {
    return {
      articleList: [],
      panigationData:{
        total:0,
        pageSize:5,
        currentPage:1,
        handleSizeChange:val=>{
          console.log(val)
        },
        handleCurrentChange:val=>{
          console.log(val)
          this.panigationData.currentPage = val
          service.get("/articles",{params:{pageNum:val}}).then(
            res=>{
              console.log(res)
              this.articleList = res.data.data.articleList
            }
          ).catch(
            err=>{
              console.log(err)
            }
          )
        }
      }
    }
  },
  methods: {
    getAllArticles: function() {
      service.get("/articles",{params:{pageNum:1}}).then(res => {
        console.log(res)
        this.articleList = res.data.data.articleList;
      });
      
    },
    dealWithTime:function(t){
      t = t.slice(0,-1)
      t = t.replace("T"," ")
      return t
    },
    showPagination(){
      service.get("/totalLength").then(
        res=>{
          // console.log(res)
          this.panigationData.total = res.data.data.total
        }
      ).catch(err=>{
        console.log(err)        
      })
    }
  },
  created() {
    // var pattern_archive  = `\\/archive\\/\\d{4}\\/\\d{1,2}` // 匹配archive路由
    if (this.$route.path.indexOf("/categories/")!=-1){
      // console.log(this.$route)
      var url = "/categories/"
      var cid = this.$route.params.cid
      service.get(url+cid)
      .then(res=>{
        // console.log(res)
        this.articleList = res.data.data.articleList
      }).catch(err=>{
        console.log(err)
        // this.$route.push("")
      })
    }else if(this.$route.path.indexOf("/archive/")!=-1){
      var archive_url = "/archive/"
      var year = this.$route.params.year
      var month = this.$route.params.month
      service.get(archive_url+year+"/"+month)
      .then(res=>{
        // console.log(res)
        this.articleList = res.data.data.articleList
      }).catch(err=>{
        console.log(err)
        // this.$route.push("")
      })
    }else{
        this.getAllArticles();
        this.showPagination()
    }
  }
};
</script>