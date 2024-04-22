/// entity服务
///
/// entity服务相关的接口
interface Interface {
    /// 新增entity
    @http.post("/v1/entities")
    create_entity(entity: Entity @1) -> Entity

    /// 获取entity详情
    @http.get("/v1/entities/{id}")
    get_entity(id: String @1) -> Entity

    /// 更新entity信息
    @http.put("/v1/entities/{id}")
    update_entity(entity: Entity @1) -> Entity

    /// 删除entity
    @http.delete("/v1/entities/{id}")
    delete_entity(id: String @1)
}
