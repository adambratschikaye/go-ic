#[no_mangle]
pub extern "C" fn init(seed: *const libc::c_uchar, seed_len: libc::size_t) {
    let seed = unsafe { std::slice::from_raw_parts(seed as *const u8, seed_len) };
    ic_wasi_polyfill::init(seed, &[("FOO", "BAR")]);
}

#[no_mangle]
pub extern "C" fn number() -> libc::c_int {
    42
}
