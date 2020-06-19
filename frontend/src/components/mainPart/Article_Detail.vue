<template>
  <article class="post post-1">
    <header class="entry-header">
      <h1 class="entry-title">{{ article_detail.Title}}</h1>
      <div class="entry-meta">
        <!-- <span class="post-category">
          <a href="#">Django 博客教程</a>
        </span>-->
        <span class="post-date">
          <a href="javascript:void(0);">
            <time
              class="entry-date"
              datetime="2012-11-09"
            >{{article_detail.CreateTime}}</time>
          </a>
        </span>
        <!-- <span class="post-author">
          <a href="#">追梦人物</a>
        </span>-->
        <span class="comments-link">
          <a href="javascript:void(0);">{{article_detail.CommentCount}} 评论</a>
        </span>
        <span class="views-count">
          <a href="javascript:void(0);">{{article_detail.ViewCount}} 阅读</a>
        </span>
      </div>
    </header>
    <div class="entry-content clearfix" v-html="htmlText"></div>
  </article>
</template>

<script>
import Marked from "marked";
import service from "@/utils/request"
export default {
  name: "article_detail",
  data() {
    return {
      article_detail: {}
    };
  },
  computed:{
    htmlText() {
      return Marked(this.article_detail.Content);
    }
  },
  methods: {
    getArticleDetail() {
      // console.log(this.$route)
      var articleId = this.$route.params.articleId;
      service
        .get("/articles/detail/" + articleId)
        .then(res => {
          this.article_detail = res.data.data.article[0];
        })
        .catch(err => {
          console.log(err);
        });
    },
  
  },
  created() {
    this.getArticleDetail();
  }
};
</script>
