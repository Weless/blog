<template>
  <div class="widget widget-category">
    <h3 class="widget-title">分类</h3>
    <ul>
      <li v-for="(category,index) in categoryList" :key="index" >
        <a :href="'/categories/'+ category.C_id">
          {{ category.Name}}
          <span class="post-count">({{ category.Cnt}})</span>
        </a>
      </li>
    </ul>
  </div>
</template>

<script>
import service from "@/utils/request"
export default {
    name:"Category",
    data(){
        return{
            categoryList:[],
        }
    },
    methods:{
        getAllCategories(){
            service.get("/categories")
            .then(res=>{
                // console.log(res)
                this.categoryList = res.data.data.categoryList
            })
        }
    },
    created(){
        this.getAllCategories()
    }
}
</script>