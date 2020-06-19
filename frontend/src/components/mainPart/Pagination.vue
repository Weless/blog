<template>
  <div class="el-pag">
    <el-pagination
    @size-change="paginationData.handleSizeChange"
    @current-change="paginationData.handleCurrentChange"
    :current-page="paginationData.currentPage"
    :page-size ="paginationData.pageSize"
    :pager-count="paginationData.pageCount"
    layout="prev, pager, next"
    :total="paginationData.total"
    ></el-pagination>
  </div>
</template>

<script>
export default {
    name:"pagination",
    props:["paginationData"]
}
</script>