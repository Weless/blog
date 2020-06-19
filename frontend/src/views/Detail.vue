<template>
  <div>
    <article-detail></article-detail>
    <!-- <comment></comment> -->
  </div>

  <!-- <div class="widget widget-content">
    <h3 class="widget-title">文章目录</h3>
    <ul>
      <li>
        <a href="#">教程特点</a>
      </li>
      <li>
        <a href="#">谁适合这个教程</a>
      </li>
      <li>
        <a href="#">在线预览</a>
      </li>
      <li>
        <a href="#">资源列表</a>
      </li>
      <li>
        <a href="#">获取帮助</a>
      </li>
    </ul>
  </div>-->
</template>

<script>
import ArticleDetail from "../components/mainPart/Article_Detail.vue";
// import Comment from "../components/mainPart/Comment.vue";

export default {
  components: {
    ArticleDetail,
    // Comment
  }
};
</script>