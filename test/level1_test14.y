
    val_CPDLQn = 12 % 3
    on (val_CPDLQn) {
        0 { pt("zero") }
        1 { pt("one") }
        2 { pt("two") }
    }
    pt(val_CPDLQn)
    

    cl Base_DTYsLa {
        pv id
        fn init(id) { this.id = id }
        fn get_id() { rev this.id }
    }
    cl Child_kZrQUV <- Base_DTYsLa {
        value
        fn init(id, val) {
            this.id = id
            this.value = val
        }
    }
    obj_aAtWYN = Child_kZrQUV(79, 38)
    pt(obj_aAtWYN.get_id())
    

    use "dummy.y"
    
    e_JryBHh = enc.b64("hello drylang")
    pt(e_JryBHh)
    
    j_jZzhxr = json(`{"test": 123}`)
    pt(j_jZzhxr)
    
    // Test get() for type info
    val_hBzzgz = get("hello")
    pt(val_hBzzgz)
    
