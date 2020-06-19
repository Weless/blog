<template>
  <div class="category-manager">
    <el-input class="title-input" v-model="newCategory" placeholder="请输入新的分类"></el-input>
    <el-button type="primary" size="small" @click="creatNewCategory">创建新的分类</el-button>

    
  </div>
</template>


<script>
import { CreateNewCategory } from "@/api/articleEditor";

export default {
  data() {
    return {
        newCategory:"",
        categoryList:[]
    };
  },
  methods: {
    creatNewCategory() {
      if (this.newCategory === "") {
        this.$message({
          message: "输入不能为空"
        });
        return;
      }
      CreateNewCategory({ name: this.newCategory })
        .then(res => {
          this.$message({
            message: res.data.msg
          });
        })
        .catch(err => {
          console.log(err);
        });
      this.getCategoryList();
    }
  }
};
</script>

<style scoped>



  .title-input {
    width: 500px;
  }

</style>