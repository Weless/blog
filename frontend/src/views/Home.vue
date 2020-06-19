<template>
  <div>
    <Nav></Nav>
    <div class="content-body">
      <div class="container">
        <div class="row">
          <main class="col-md-8">
            <router-view />
          </main>
          <aside class="col-md-4">
            <Latest></Latest>
            <Archive></Archive>
            <Category></Category>
          </aside>
        </div>
      </div>
    </div>
    <Footer></Footer>
  </div>
</template>

<script>
// @ is an alias to /src
// import Article from "../components/mainPart/Article.vue";
import Nav from "../components/publicPart/Nav"
import Footer from "../components/publicPart/Footer"
import Latest from "../components/sidePart/Latest"
import Archive from "../components/sidePart/Archive"
import Category from "../components/sidePart/Category"

export default {
  name: "Home",
  components:{
    Nav,
    Footer,
    Latest,
    Archive,
    Category,
  }
};
</script> 

