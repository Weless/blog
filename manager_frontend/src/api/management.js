import service from "@/utils/request";

export const GetArticleList = data => {
  return service.request({
    method: "get",
    url: "/articleList",
    params: data
  });
};

// export const ChangeState = data => {
//   return service.request({
//     method: "post",
//     url: "/api/article/change",
//     data
//   });
// };

export const RemoveArticle = data => {
  return service.request({
    method: "post",
    url: "/deleteArticle",
    data
  });
};
