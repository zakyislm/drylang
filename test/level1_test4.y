
    arr_joMxhP = [74, 5+1, 96+2]
    map_YgyxUJ = {"a": arr_joMxhP[0], "b": arr_joMxhP[1]}
    pt(map_YgyxUJ["a"])
    

    cl Base_ubxQhH {
        pv id
        fn init(id) { this.id = id }
        fn get_id() { rev this.id }
    }
    cl Child_VsQLwS <- Base_ubxQhH {
        value
        fn init(id, val) {
            this.id = id
            this.value = val
        }
    }
    obj_Yclscz = Child_VsQLwS(1, 94)
    pt(obj_Yclscz.get_id())
    

    cns C_vhIxeF = 24
    v_ToNZrn = unknown
    v_ToNZrn = C_vhIxeF + 20
    if (v_ToNZrn > 0) {
        v_ToNZrn = v_ToNZrn * 2
    } el {
        v_ToNZrn = 0
    }
    pt(v_ToNZrn)
    
