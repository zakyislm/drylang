
    cl Base_GQiYhz {
        pv id
        fn init(id) { this.id = id }
        fn get_id() { rev this.id }
    }
    cl Child_wUDEPO <- Base_GQiYhz {
        value
        fn init(id, val) {
            this.id = id
            this.value = val
        }
    }
    obj_gsfRCd = Child_wUDEPO(84, 67)
    pt(obj_gsfRCd.get_id())
    

    m_fUkCYV = math.sqrt(2)
    h_JqTQCm = hash.md5("test_NMAKsp")
    log.inf("hash:", h_JqTQCm)
    pt(m_fUkCYV)
    

    use "dummy.y"
    
    e_iWZExb = enc.b64("hello drylang")
    pt(e_iWZExb)
    
    j_CTWUaX = json(`{"test": 123}`)
    pt(j_CTWUaX)
    
    // Test get() for type info
    val_QdUggT = get("hello")
    pt(val_QdUggT)
    
