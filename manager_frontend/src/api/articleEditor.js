import service from "@/utils/request";

export const SaveArticle = data => {
  return service.request({
    method: "post",
    url: "/insertArticle",
    data
  });
};

export const ChangeArticle = data => {
  return service.request({
    method: "post",
    url: "/updateArticle",
    data
  });
};

export const GetArticle = data => {
  return service.request({
    method: "get",
    url: "/articleDetail",
    params: data
  });
};

export const GetCategoryList = () =>{
  return service.request({
    method:"get",
    url:"/categoryList",
  })
}

export const CreateNewCategory = data =>{
  return service.request({
    method:"post",
    url:"insertCategory",
    data
  })
}