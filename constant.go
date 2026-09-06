// This is a generated file. DO NOT EDIT.

package glfw

/*
The major version number of the GLFW header.  This is incremented when the
API is changed in non-compatible ways.

C documentation : [GLFW_VERSION_MAJOR]

[GLFW_VERSION_MAJOR]: https://www.glfw.org/docs/latest/group__init.html#ga6337d9ea43b22fc529b2bba066b4a576
*/
const VERSION_MAJOR = 3

/*
The minor version number of the GLFW header.  This is incremented when
features are added to the API but it remains backward-compatible.

C documentation : [GLFW_VERSION_MINOR]

[GLFW_VERSION_MINOR]: https://www.glfw.org/docs/latest/group__init.html#gaf80d40f0aea7088ff337606e9c48f7a3
*/
const VERSION_MINOR = 4

/*
The revision number of the GLFW header.  This is incremented when a bug fix
release is made that does not contain any API changes.

C documentation : [GLFW_VERSION_REVISION]

[GLFW_VERSION_REVISION]: https://www.glfw.org/docs/latest/group__init.html#gab72ae2e2035d9ea461abc3495eac0502
*/
const VERSION_REVISION = 0

/*
This is only semantic sugar for the number 1.  You can instead use `1` or
`true` or `_True` or `GL_TRUE` or `VK_TRUE` or anything else that is equal
to one.

C documentation : [GLFW_TRUE]

[GLFW_TRUE]: https://www.glfw.org/docs/latest/group__init.html#ga2744fbb29b5631bb28802dbe0cf36eba
*/
const TRUE = 1

/*
This is only semantic sugar for the number 0.  You can instead use `0` or
`false` or `_False` or `GL_FALSE` or `VK_FALSE` or anything else that is
equal to zero.

C documentation : [GLFW_FALSE]

[GLFW_FALSE]: https://www.glfw.org/docs/latest/group__init.html#gac877fe3b627d21ef3a0a23e0a73ba8c5
*/
const FALSE = 0

/*
The key or mouse button was released.

C documentation : [GLFW_RELEASE]

[GLFW_RELEASE]: https://www.glfw.org/docs/latest/group__input.html#gada11d965c4da13090ad336e030e4d11f
*/
const RELEASE = 0

/*
The key or mouse button was pressed.

C documentation : [GLFW_PRESS]

[GLFW_PRESS]: https://www.glfw.org/docs/latest/group__input.html#ga2485743d0b59df3791c45951c4195265
*/
const PRESS = 1

/*
The key was held down until it repeated.

C documentation : [GLFW_REPEAT]

[GLFW_REPEAT]: https://www.glfw.org/docs/latest/group__input.html#gac96fd3b9fc66c6f0eebaf6532595338f
*/
const REPEAT = 2

/*
See [joystick hat input] for how these are used.

C documentation : [GLFW_HAT_CENTERED]

[joystick hat input]: https://www.glfw.org/docs/latest/input_guide.html#joystick_hat
[GLFW_HAT_CENTERED]: https://www.glfw.org/docs/latest/group__hat__state.html#gae2c0bcb7aec609e4736437554f6638fd
*/
const HAT_CENTERED = 0

/*
C documentation : [GLFW_HAT_UP]

[GLFW_HAT_UP]: https://www.glfw.org/docs/latest/group__hat__state.html#ga8c9720c76cd1b912738159ed74c85b36
*/
const HAT_UP = 1

/*
C documentation : [GLFW_HAT_RIGHT]

[GLFW_HAT_RIGHT]: https://www.glfw.org/docs/latest/group__hat__state.html#ga252586e3bbde75f4b0e07ad3124867f5
*/
const HAT_RIGHT = 2

/*
C documentation : [GLFW_HAT_DOWN]

[GLFW_HAT_DOWN]: https://www.glfw.org/docs/latest/group__hat__state.html#gad60d1fd0dc85c18f2642cbae96d3deff
*/
const HAT_DOWN = 4

/*
C documentation : [GLFW_HAT_LEFT]

[GLFW_HAT_LEFT]: https://www.glfw.org/docs/latest/group__hat__state.html#gac775f4b3154fdf5db93eb432ba546dff
*/
const HAT_LEFT = 8

/*
C documentation : [GLFW_HAT_RIGHT_UP]

[GLFW_HAT_RIGHT_UP]: https://www.glfw.org/docs/latest/group__hat__state.html#ga94aea0ae241a8b902883536c592ee693
*/
const HAT_RIGHT_UP = (HAT_RIGHT | HAT_UP)

/*
C documentation : [GLFW_HAT_RIGHT_DOWN]

[GLFW_HAT_RIGHT_DOWN]: https://www.glfw.org/docs/latest/group__hat__state.html#gad7f0e4f52fd68d734863aaeadab3a3f5
*/
const HAT_RIGHT_DOWN = (HAT_RIGHT | HAT_DOWN)

/*
C documentation : [GLFW_HAT_LEFT_UP]

[GLFW_HAT_LEFT_UP]: https://www.glfw.org/docs/latest/group__hat__state.html#ga638f0e20dc5de90de21a33564e8ce129
*/
const HAT_LEFT_UP = (HAT_LEFT | HAT_UP)

/*
C documentation : [GLFW_HAT_LEFT_DOWN]

[GLFW_HAT_LEFT_DOWN]: https://www.glfw.org/docs/latest/group__hat__state.html#ga76c02baf1ea345fcbe3e8ff176a73e19
*/
const HAT_LEFT_DOWN = (HAT_LEFT | HAT_DOWN)

/*
C documentation : [GLFW_KEY_UNKNOWN]

[GLFW_KEY_UNKNOWN]: https://www.glfw.org/docs/latest/group__input.html#ga99aacc875b6b27a072552631e13775c7
*/
const KEY_UNKNOWN = -1

/*
C documentation : [GLFW_KEY_SPACE]

[GLFW_KEY_SPACE]: https://www.glfw.org/docs/latest/group__keys.html#gaddb2c23772b97fd7e26e8ee66f1ad014
*/
const KEY_SPACE = 32

/*
C documentation : [GLFW_KEY_APOSTROPHE]

[GLFW_KEY_APOSTROPHE]: https://www.glfw.org/docs/latest/group__keys.html#ga6059b0b048ba6980b6107fffbd3b4b24
*/
const KEY_APOSTROPHE = 39

/*
C documentation : [GLFW_KEY_COMMA]

[GLFW_KEY_COMMA]: https://www.glfw.org/docs/latest/group__keys.html#gab3d5d72e59d3055f494627b0a524926c
*/
const KEY_COMMA = 44

/*
C documentation : [GLFW_KEY_MINUS]

[GLFW_KEY_MINUS]: https://www.glfw.org/docs/latest/group__keys.html#gac556b360f7f6fca4b70ba0aecf313fd4
*/
const KEY_MINUS = 45

/*
C documentation : [GLFW_KEY_PERIOD]

[GLFW_KEY_PERIOD]: https://www.glfw.org/docs/latest/group__keys.html#ga37e296b650eab419fc474ff69033d927
*/
const KEY_PERIOD = 46

/*
C documentation : [GLFW_KEY_SLASH]

[GLFW_KEY_SLASH]: https://www.glfw.org/docs/latest/group__keys.html#gadf3d753b2d479148d711de34b83fd0db
*/
const KEY_SLASH = 47

/*
C documentation : [GLFW_KEY_0]

[GLFW_KEY_0]: https://www.glfw.org/docs/latest/group__keys.html#ga50391730e9d7112ad4fd42d0bd1597c1
*/
const KEY_0 = 48

/*
C documentation : [GLFW_KEY_1]

[GLFW_KEY_1]: https://www.glfw.org/docs/latest/group__keys.html#ga05e4cae9ddb8d40cf6d82c8f11f2502f
*/
const KEY_1 = 49

/*
C documentation : [GLFW_KEY_2]

[GLFW_KEY_2]: https://www.glfw.org/docs/latest/group__keys.html#gadc8e66b3a4c4b5c39ad1305cf852863c
*/
const KEY_2 = 50

/*
C documentation : [GLFW_KEY_3]

[GLFW_KEY_3]: https://www.glfw.org/docs/latest/group__keys.html#ga812f0273fe1a981e1fa002ae73e92271
*/
const KEY_3 = 51

/*
C documentation : [GLFW_KEY_4]

[GLFW_KEY_4]: https://www.glfw.org/docs/latest/group__keys.html#ga9e14b6975a9cc8f66cdd5cb3d3861356
*/
const KEY_4 = 52

/*
C documentation : [GLFW_KEY_5]

[GLFW_KEY_5]: https://www.glfw.org/docs/latest/group__keys.html#ga4d74ddaa5d4c609993b4d4a15736c924
*/
const KEY_5 = 53

/*
C documentation : [GLFW_KEY_6]

[GLFW_KEY_6]: https://www.glfw.org/docs/latest/group__keys.html#ga9ea4ab80c313a227b14d0a7c6f810b5d
*/
const KEY_6 = 54

/*
C documentation : [GLFW_KEY_7]

[GLFW_KEY_7]: https://www.glfw.org/docs/latest/group__keys.html#gab79b1cfae7bd630cfc4604c1f263c666
*/
const KEY_7 = 55

/*
C documentation : [GLFW_KEY_8]

[GLFW_KEY_8]: https://www.glfw.org/docs/latest/group__keys.html#gadeaa109a0f9f5afc94fe4a108e686f6f
*/
const KEY_8 = 56

/*
C documentation : [GLFW_KEY_9]

[GLFW_KEY_9]: https://www.glfw.org/docs/latest/group__keys.html#ga2924cb5349ebbf97c8987f3521c44f39
*/
const KEY_9 = 57

/*
C documentation : [GLFW_KEY_SEMICOLON]

[GLFW_KEY_SEMICOLON]: https://www.glfw.org/docs/latest/group__keys.html#ga84233de9ee5bb3e8788a5aa07d80af7d
*/
const KEY_SEMICOLON = 59

/*
C documentation : [GLFW_KEY_EQUAL]

[GLFW_KEY_EQUAL]: https://www.glfw.org/docs/latest/group__keys.html#gae1a2de47240d6664423c204bdd91bd17
*/
const KEY_EQUAL = 61

/*
C documentation : [GLFW_KEY_A]

[GLFW_KEY_A]: https://www.glfw.org/docs/latest/group__keys.html#ga03e842608e1ea323370889d33b8f70ff
*/
const KEY_A = 65

/*
C documentation : [GLFW_KEY_B]

[GLFW_KEY_B]: https://www.glfw.org/docs/latest/group__keys.html#ga8e3fb647ff3aca9e8dbf14fe66332941
*/
const KEY_B = 66

/*
C documentation : [GLFW_KEY_C]

[GLFW_KEY_C]: https://www.glfw.org/docs/latest/group__keys.html#ga00ccf3475d9ee2e679480d540d554669
*/
const KEY_C = 67

/*
C documentation : [GLFW_KEY_D]

[GLFW_KEY_D]: https://www.glfw.org/docs/latest/group__keys.html#ga011f7cdc9a654da984a2506479606933
*/
const KEY_D = 68

/*
C documentation : [GLFW_KEY_E]

[GLFW_KEY_E]: https://www.glfw.org/docs/latest/group__keys.html#gabf48fcc3afbe69349df432b470c96ef2
*/
const KEY_E = 69

/*
C documentation : [GLFW_KEY_F]

[GLFW_KEY_F]: https://www.glfw.org/docs/latest/group__keys.html#ga5df402e02aca08444240058fd9b42a55
*/
const KEY_F = 70

/*
C documentation : [GLFW_KEY_G]

[GLFW_KEY_G]: https://www.glfw.org/docs/latest/group__keys.html#gae74ecddf7cc96104ab23989b1cdab536
*/
const KEY_G = 71

/*
C documentation : [GLFW_KEY_H]

[GLFW_KEY_H]: https://www.glfw.org/docs/latest/group__keys.html#gad4cc98fc8f35f015d9e2fb94bf136076
*/
const KEY_H = 72

/*
C documentation : [GLFW_KEY_I]

[GLFW_KEY_I]: https://www.glfw.org/docs/latest/group__keys.html#ga274655c8bfe39742684ca393cf8ed093
*/
const KEY_I = 73

/*
C documentation : [GLFW_KEY_J]

[GLFW_KEY_J]: https://www.glfw.org/docs/latest/group__keys.html#ga65ff2aedb129a3149ad9cb3e4159a75f
*/
const KEY_J = 74

/*
C documentation : [GLFW_KEY_K]

[GLFW_KEY_K]: https://www.glfw.org/docs/latest/group__keys.html#ga4ae8debadf6d2a691badae0b53ea3ba0
*/
const KEY_K = 75

/*
C documentation : [GLFW_KEY_L]

[GLFW_KEY_L]: https://www.glfw.org/docs/latest/group__keys.html#gaaa8b54a13f6b1eed85ac86f82d550db2
*/
const KEY_L = 76

/*
C documentation : [GLFW_KEY_M]

[GLFW_KEY_M]: https://www.glfw.org/docs/latest/group__keys.html#ga4d7f0260c82e4ea3d6ebc7a21d6e3716
*/
const KEY_M = 77

/*
C documentation : [GLFW_KEY_N]

[GLFW_KEY_N]: https://www.glfw.org/docs/latest/group__keys.html#gae00856dfeb5d13aafebf59d44de5cdda
*/
const KEY_N = 78

/*
C documentation : [GLFW_KEY_O]

[GLFW_KEY_O]: https://www.glfw.org/docs/latest/group__keys.html#gaecbbb79130df419d58dd7f09a169efe9
*/
const KEY_O = 79

/*
C documentation : [GLFW_KEY_P]

[GLFW_KEY_P]: https://www.glfw.org/docs/latest/group__keys.html#ga8fc15819c1094fb2afa01d84546b33e1
*/
const KEY_P = 80

/*
C documentation : [GLFW_KEY_Q]

[GLFW_KEY_Q]: https://www.glfw.org/docs/latest/group__keys.html#gafdd01e38b120d67cf51e348bb47f3964
*/
const KEY_Q = 81

/*
C documentation : [GLFW_KEY_R]

[GLFW_KEY_R]: https://www.glfw.org/docs/latest/group__keys.html#ga4ce6c70a0c98c50b3fe4ab9a728d4d36
*/
const KEY_R = 82

/*
C documentation : [GLFW_KEY_S]

[GLFW_KEY_S]: https://www.glfw.org/docs/latest/group__keys.html#ga1570e2ccaab036ea82bed66fc1dab2a9
*/
const KEY_S = 83

/*
C documentation : [GLFW_KEY_T]

[GLFW_KEY_T]: https://www.glfw.org/docs/latest/group__keys.html#ga90e0560422ec7a30e7f3f375bc9f37f9
*/
const KEY_T = 84

/*
C documentation : [GLFW_KEY_U]

[GLFW_KEY_U]: https://www.glfw.org/docs/latest/group__keys.html#gacad52f3bf7d378fc0ffa72a76769256d
*/
const KEY_U = 85

/*
C documentation : [GLFW_KEY_V]

[GLFW_KEY_V]: https://www.glfw.org/docs/latest/group__keys.html#ga22c7763899ecf7788862e5f90eacce6b
*/
const KEY_V = 86

/*
C documentation : [GLFW_KEY_W]

[GLFW_KEY_W]: https://www.glfw.org/docs/latest/group__keys.html#gaa06a712e6202661fc03da5bdb7b6e545
*/
const KEY_W = 87

/*
C documentation : [GLFW_KEY_X]

[GLFW_KEY_X]: https://www.glfw.org/docs/latest/group__keys.html#gac1c42c0bf4192cea713c55598b06b744
*/
const KEY_X = 88

/*
C documentation : [GLFW_KEY_Y]

[GLFW_KEY_Y]: https://www.glfw.org/docs/latest/group__keys.html#gafd9f115a549effdf8e372a787c360313
*/
const KEY_Y = 89

/*
C documentation : [GLFW_KEY_Z]

[GLFW_KEY_Z]: https://www.glfw.org/docs/latest/group__keys.html#gac489e208c26afda8d4938ed88718760a
*/
const KEY_Z = 90

/*
C documentation : [GLFW_KEY_LEFT_BRACKET]

[GLFW_KEY_LEFT_BRACKET]: https://www.glfw.org/docs/latest/group__keys.html#gad1c8d9adac53925276ecb1d592511d8a
*/
const KEY_LEFT_BRACKET = 91

/*
C documentation : [GLFW_KEY_BACKSLASH]

[GLFW_KEY_BACKSLASH]: https://www.glfw.org/docs/latest/group__keys.html#gab8155ea99d1ab27ff56f24f8dc73f8d1
*/
const KEY_BACKSLASH = 92

/*
C documentation : [GLFW_KEY_RIGHT_BRACKET]

[GLFW_KEY_RIGHT_BRACKET]: https://www.glfw.org/docs/latest/group__keys.html#ga86ef225fd6a66404caae71044cdd58d8
*/
const KEY_RIGHT_BRACKET = 93

/*
C documentation : [GLFW_KEY_GRAVE_ACCENT]

[GLFW_KEY_GRAVE_ACCENT]: https://www.glfw.org/docs/latest/group__keys.html#ga7a3701fb4e2a0b136ff4b568c3c8d668
*/
const KEY_GRAVE_ACCENT = 96

/*
C documentation : [GLFW_KEY_WORLD_1]

[GLFW_KEY_WORLD_1]: https://www.glfw.org/docs/latest/group__keys.html#gadc78dad3dab76bcd4b5c20114052577a
*/
const KEY_WORLD_1 = 161

/*
C documentation : [GLFW_KEY_WORLD_2]

[GLFW_KEY_WORLD_2]: https://www.glfw.org/docs/latest/group__keys.html#ga20494bfebf0bb4fc9503afca18ab2c5e
*/
const KEY_WORLD_2 = 162

/*
C documentation : [GLFW_KEY_ESCAPE]

[GLFW_KEY_ESCAPE]: https://www.glfw.org/docs/latest/group__keys.html#gaac6596c350b635c245113b81c2123b93
*/
const KEY_ESCAPE = 256

/*
C documentation : [GLFW_KEY_ENTER]

[GLFW_KEY_ENTER]: https://www.glfw.org/docs/latest/group__keys.html#ga9555a92ecbecdbc1f3435219c571d667
*/
const KEY_ENTER = 257

/*
C documentation : [GLFW_KEY_TAB]

[GLFW_KEY_TAB]: https://www.glfw.org/docs/latest/group__keys.html#ga6908a4bda9950a3e2b73f794bbe985df
*/
const KEY_TAB = 258

/*
C documentation : [GLFW_KEY_BACKSPACE]

[GLFW_KEY_BACKSPACE]: https://www.glfw.org/docs/latest/group__keys.html#ga6c0df1fe2f156bbd5a98c66d76ff3635
*/
const KEY_BACKSPACE = 259

/*
C documentation : [GLFW_KEY_INSERT]

[GLFW_KEY_INSERT]: https://www.glfw.org/docs/latest/group__keys.html#ga373ac7365435d6b0eb1068f470e34f47
*/
const KEY_INSERT = 260

/*
C documentation : [GLFW_KEY_DELETE]

[GLFW_KEY_DELETE]: https://www.glfw.org/docs/latest/group__keys.html#gadb111e4df74b8a715f2c05dad58d2682
*/
const KEY_DELETE = 261

/*
C documentation : [GLFW_KEY_RIGHT]

[GLFW_KEY_RIGHT]: https://www.glfw.org/docs/latest/group__keys.html#ga06ba07662e8c291a4a84535379ffc7ac
*/
const KEY_RIGHT = 262

/*
C documentation : [GLFW_KEY_LEFT]

[GLFW_KEY_LEFT]: https://www.glfw.org/docs/latest/group__keys.html#gae12a010d33c309a67ab9460c51eb2462
*/
const KEY_LEFT = 263

/*
C documentation : [GLFW_KEY_DOWN]

[GLFW_KEY_DOWN]: https://www.glfw.org/docs/latest/group__keys.html#gae2e3958c71595607416aa7bf082be2f9
*/
const KEY_DOWN = 264

/*
C documentation : [GLFW_KEY_UP]

[GLFW_KEY_UP]: https://www.glfw.org/docs/latest/group__keys.html#ga2f3342b194020d3544c67e3506b6f144
*/
const KEY_UP = 265

/*
C documentation : [GLFW_KEY_PAGE_UP]

[GLFW_KEY_PAGE_UP]: https://www.glfw.org/docs/latest/group__keys.html#ga3ab731f9622f0db280178a5f3cc6d586
*/
const KEY_PAGE_UP = 266

/*
C documentation : [GLFW_KEY_PAGE_DOWN]

[GLFW_KEY_PAGE_DOWN]: https://www.glfw.org/docs/latest/group__keys.html#gaee0a8fa442001cc2147812f84b59041c
*/
const KEY_PAGE_DOWN = 267

/*
C documentation : [GLFW_KEY_HOME]

[GLFW_KEY_HOME]: https://www.glfw.org/docs/latest/group__keys.html#ga41452c7287195d481e43207318c126a7
*/
const KEY_HOME = 268

/*
C documentation : [GLFW_KEY_END]

[GLFW_KEY_END]: https://www.glfw.org/docs/latest/group__keys.html#ga86587ea1df19a65978d3e3b8439bedd9
*/
const KEY_END = 269

/*
C documentation : [GLFW_KEY_CAPS_LOCK]

[GLFW_KEY_CAPS_LOCK]: https://www.glfw.org/docs/latest/group__keys.html#ga92c1d2c9d63485f3d70f94f688d48672
*/
const KEY_CAPS_LOCK = 280

/*
C documentation : [GLFW_KEY_SCROLL_LOCK]

[GLFW_KEY_SCROLL_LOCK]: https://www.glfw.org/docs/latest/group__keys.html#gaf622b63b9537f7084c2ab649b8365630
*/
const KEY_SCROLL_LOCK = 281

/*
C documentation : [GLFW_KEY_NUM_LOCK]

[GLFW_KEY_NUM_LOCK]: https://www.glfw.org/docs/latest/group__keys.html#ga3946edc362aeff213b2be6304296cf43
*/
const KEY_NUM_LOCK = 282

/*
C documentation : [GLFW_KEY_PRINT_SCREEN]

[GLFW_KEY_PRINT_SCREEN]: https://www.glfw.org/docs/latest/group__keys.html#gaf964c2e65e97d0cf785a5636ee8df642
*/
const KEY_PRINT_SCREEN = 283

/*
C documentation : [GLFW_KEY_PAUSE]

[GLFW_KEY_PAUSE]: https://www.glfw.org/docs/latest/group__keys.html#ga8116b9692d87382afb5849b6d8907f18
*/
const KEY_PAUSE = 284

/*
C documentation : [GLFW_KEY_F1]

[GLFW_KEY_F1]: https://www.glfw.org/docs/latest/group__keys.html#gafb8d66c573acf22e364049477dcbea30
*/
const KEY_F1 = 290

/*
C documentation : [GLFW_KEY_F2]

[GLFW_KEY_F2]: https://www.glfw.org/docs/latest/group__keys.html#ga0900750aff94889b940f5e428c07daee
*/
const KEY_F2 = 291

/*
C documentation : [GLFW_KEY_F3]

[GLFW_KEY_F3]: https://www.glfw.org/docs/latest/group__keys.html#gaed7cd729c0147a551bb8b7bb36c17015
*/
const KEY_F3 = 292

/*
C documentation : [GLFW_KEY_F4]

[GLFW_KEY_F4]: https://www.glfw.org/docs/latest/group__keys.html#ga9b61ebd0c63b44b7332fda2c9763eaa6
*/
const KEY_F4 = 293

/*
C documentation : [GLFW_KEY_F5]

[GLFW_KEY_F5]: https://www.glfw.org/docs/latest/group__keys.html#gaf258dda9947daa428377938ed577c8c2
*/
const KEY_F5 = 294

/*
C documentation : [GLFW_KEY_F6]

[GLFW_KEY_F6]: https://www.glfw.org/docs/latest/group__keys.html#ga6dc2d3f87b9d51ffbbbe2ef0299d8e1d
*/
const KEY_F6 = 295

/*
C documentation : [GLFW_KEY_F7]

[GLFW_KEY_F7]: https://www.glfw.org/docs/latest/group__keys.html#gacca6ef8a2162c52a0ac1d881e8d9c38a
*/
const KEY_F7 = 296

/*
C documentation : [GLFW_KEY_F8]

[GLFW_KEY_F8]: https://www.glfw.org/docs/latest/group__keys.html#gac9d39390336ae14e4a93e295de43c7e8
*/
const KEY_F8 = 297

/*
C documentation : [GLFW_KEY_F9]

[GLFW_KEY_F9]: https://www.glfw.org/docs/latest/group__keys.html#gae40de0de1c9f21cd26c9afa3d7050851
*/
const KEY_F9 = 298

/*
C documentation : [GLFW_KEY_F10]

[GLFW_KEY_F10]: https://www.glfw.org/docs/latest/group__keys.html#ga718d11d2f7d57471a2f6a894235995b1
*/
const KEY_F10 = 299

/*
C documentation : [GLFW_KEY_F11]

[GLFW_KEY_F11]: https://www.glfw.org/docs/latest/group__keys.html#ga0bc04b11627e7d69339151e7306b2832
*/
const KEY_F11 = 300

/*
C documentation : [GLFW_KEY_F12]

[GLFW_KEY_F12]: https://www.glfw.org/docs/latest/group__keys.html#gaf5908fa9b0a906ae03fc2c61ac7aa3e2
*/
const KEY_F12 = 301

/*
C documentation : [GLFW_KEY_F13]

[GLFW_KEY_F13]: https://www.glfw.org/docs/latest/group__keys.html#gad637f4308655e1001bd6ad942bc0fd4b
*/
const KEY_F13 = 302

/*
C documentation : [GLFW_KEY_F14]

[GLFW_KEY_F14]: https://www.glfw.org/docs/latest/group__keys.html#gaf14c66cff3396e5bd46e803c035e6c1f
*/
const KEY_F14 = 303

/*
C documentation : [GLFW_KEY_F15]

[GLFW_KEY_F15]: https://www.glfw.org/docs/latest/group__keys.html#ga7f70970db6e8be1794da8516a6d14058
*/
const KEY_F15 = 304

/*
C documentation : [GLFW_KEY_F16]

[GLFW_KEY_F16]: https://www.glfw.org/docs/latest/group__keys.html#gaa582dbb1d2ba2050aa1dca0838095b27
*/
const KEY_F16 = 305

/*
C documentation : [GLFW_KEY_F17]

[GLFW_KEY_F17]: https://www.glfw.org/docs/latest/group__keys.html#ga972ce5c365e2394b36104b0e3125c748
*/
const KEY_F17 = 306

/*
C documentation : [GLFW_KEY_F18]

[GLFW_KEY_F18]: https://www.glfw.org/docs/latest/group__keys.html#gaebf6391058d5566601e357edc5ea737c
*/
const KEY_F18 = 307

/*
C documentation : [GLFW_KEY_F19]

[GLFW_KEY_F19]: https://www.glfw.org/docs/latest/group__keys.html#gaec011d9ba044058cb54529da710e9791
*/
const KEY_F19 = 308

/*
C documentation : [GLFW_KEY_F20]

[GLFW_KEY_F20]: https://www.glfw.org/docs/latest/group__keys.html#ga82b9c721ada04cd5ca8de767da38022f
*/
const KEY_F20 = 309

/*
C documentation : [GLFW_KEY_F21]

[GLFW_KEY_F21]: https://www.glfw.org/docs/latest/group__keys.html#ga356afb14d3440ff2bb378f74f7ebc60f
*/
const KEY_F21 = 310

/*
C documentation : [GLFW_KEY_F22]

[GLFW_KEY_F22]: https://www.glfw.org/docs/latest/group__keys.html#ga90960bd2a155f2b09675324d3dff1565
*/
const KEY_F22 = 311

/*
C documentation : [GLFW_KEY_F23]

[GLFW_KEY_F23]: https://www.glfw.org/docs/latest/group__keys.html#ga43c21099aac10952d1be909a8ddee4d5
*/
const KEY_F23 = 312

/*
C documentation : [GLFW_KEY_F24]

[GLFW_KEY_F24]: https://www.glfw.org/docs/latest/group__keys.html#ga8150374677b5bed3043408732152dea2
*/
const KEY_F24 = 313

/*
C documentation : [GLFW_KEY_F25]

[GLFW_KEY_F25]: https://www.glfw.org/docs/latest/group__keys.html#gaa4bbd93ed73bb4c6ae7d83df880b7199
*/
const KEY_F25 = 314

/*
C documentation : [GLFW_KEY_KP_0]

[GLFW_KEY_KP_0]: https://www.glfw.org/docs/latest/group__keys.html#ga10515dafc55b71e7683f5b4fedd1c70d
*/
const KEY_KP_0 = 320

/*
C documentation : [GLFW_KEY_KP_1]

[GLFW_KEY_KP_1]: https://www.glfw.org/docs/latest/group__keys.html#gaf3a29a334402c5eaf0b3439edf5587c3
*/
const KEY_KP_1 = 321

/*
C documentation : [GLFW_KEY_KP_2]

[GLFW_KEY_KP_2]: https://www.glfw.org/docs/latest/group__keys.html#gaf82d5a802ab8213c72653d7480c16f13
*/
const KEY_KP_2 = 322

/*
C documentation : [GLFW_KEY_KP_3]

[GLFW_KEY_KP_3]: https://www.glfw.org/docs/latest/group__keys.html#ga7e25ff30d56cd512828c1d4ae8d54ef2
*/
const KEY_KP_3 = 323

/*
C documentation : [GLFW_KEY_KP_4]

[GLFW_KEY_KP_4]: https://www.glfw.org/docs/latest/group__keys.html#gada7ec86778b85e0b4de0beea72234aea
*/
const KEY_KP_4 = 324

/*
C documentation : [GLFW_KEY_KP_5]

[GLFW_KEY_KP_5]: https://www.glfw.org/docs/latest/group__keys.html#ga9a5be274434866c51738cafbb6d26b45
*/
const KEY_KP_5 = 325

/*
C documentation : [GLFW_KEY_KP_6]

[GLFW_KEY_KP_6]: https://www.glfw.org/docs/latest/group__keys.html#gafc141b0f8450519084c01092a3157faa
*/
const KEY_KP_6 = 326

/*
C documentation : [GLFW_KEY_KP_7]

[GLFW_KEY_KP_7]: https://www.glfw.org/docs/latest/group__keys.html#ga8882f411f05d04ec77a9563974bbfa53
*/
const KEY_KP_7 = 327

/*
C documentation : [GLFW_KEY_KP_8]

[GLFW_KEY_KP_8]: https://www.glfw.org/docs/latest/group__keys.html#gab2ea2e6a12f89d315045af520ac78cec
*/
const KEY_KP_8 = 328

/*
C documentation : [GLFW_KEY_KP_9]

[GLFW_KEY_KP_9]: https://www.glfw.org/docs/latest/group__keys.html#gafb21426b630ed4fcc084868699ba74c1
*/
const KEY_KP_9 = 329

/*
C documentation : [GLFW_KEY_KP_DECIMAL]

[GLFW_KEY_KP_DECIMAL]: https://www.glfw.org/docs/latest/group__keys.html#ga4e231d968796331a9ea0dbfb98d4005b
*/
const KEY_KP_DECIMAL = 330

/*
C documentation : [GLFW_KEY_KP_DIVIDE]

[GLFW_KEY_KP_DIVIDE]: https://www.glfw.org/docs/latest/group__keys.html#gabca1733780a273d549129ad0f250d1e5
*/
const KEY_KP_DIVIDE = 331

/*
C documentation : [GLFW_KEY_KP_MULTIPLY]

[GLFW_KEY_KP_MULTIPLY]: https://www.glfw.org/docs/latest/group__keys.html#ga9ada267eb0e78ed2ada8701dd24a56ef
*/
const KEY_KP_MULTIPLY = 332

/*
C documentation : [GLFW_KEY_KP_SUBTRACT]

[GLFW_KEY_KP_SUBTRACT]: https://www.glfw.org/docs/latest/group__keys.html#gaa3dbd60782ff93d6082a124bce1fa236
*/
const KEY_KP_SUBTRACT = 333

/*
C documentation : [GLFW_KEY_KP_ADD]

[GLFW_KEY_KP_ADD]: https://www.glfw.org/docs/latest/group__keys.html#gad09c7c98acc79e89aa6a0a91275becac
*/
const KEY_KP_ADD = 334

/*
C documentation : [GLFW_KEY_KP_ENTER]

[GLFW_KEY_KP_ENTER]: https://www.glfw.org/docs/latest/group__keys.html#ga4f728f8738f2986bd63eedd3d412e8cf
*/
const KEY_KP_ENTER = 335

/*
C documentation : [GLFW_KEY_KP_EQUAL]

[GLFW_KEY_KP_EQUAL]: https://www.glfw.org/docs/latest/group__keys.html#gaebdc76d4a808191e6d21b7e4ad2acd97
*/
const KEY_KP_EQUAL = 336

/*
C documentation : [GLFW_KEY_LEFT_SHIFT]

[GLFW_KEY_LEFT_SHIFT]: https://www.glfw.org/docs/latest/group__keys.html#ga8a530a28a65c44ab5d00b759b756d3f6
*/
const KEY_LEFT_SHIFT = 340

/*
C documentation : [GLFW_KEY_LEFT_CONTROL]

[GLFW_KEY_LEFT_CONTROL]: https://www.glfw.org/docs/latest/group__keys.html#ga9f97b743e81460ac4b2deddecd10a464
*/
const KEY_LEFT_CONTROL = 341

/*
C documentation : [GLFW_KEY_LEFT_ALT]

[GLFW_KEY_LEFT_ALT]: https://www.glfw.org/docs/latest/group__keys.html#ga7f27dabf63a7789daa31e1c96790219b
*/
const KEY_LEFT_ALT = 342

/*
C documentation : [GLFW_KEY_LEFT_SUPER]

[GLFW_KEY_LEFT_SUPER]: https://www.glfw.org/docs/latest/group__keys.html#gafb1207c91997fc295afd1835fbc5641a
*/
const KEY_LEFT_SUPER = 343

/*
C documentation : [GLFW_KEY_RIGHT_SHIFT]

[GLFW_KEY_RIGHT_SHIFT]: https://www.glfw.org/docs/latest/group__keys.html#gaffca36b99c9dce1a19cb9befbadce691
*/
const KEY_RIGHT_SHIFT = 344

/*
C documentation : [GLFW_KEY_RIGHT_CONTROL]

[GLFW_KEY_RIGHT_CONTROL]: https://www.glfw.org/docs/latest/group__keys.html#gad1ca2094b2694e7251d0ab1fd34f8519
*/
const KEY_RIGHT_CONTROL = 345

/*
C documentation : [GLFW_KEY_RIGHT_ALT]

[GLFW_KEY_RIGHT_ALT]: https://www.glfw.org/docs/latest/group__keys.html#ga687b38009131cfdd07a8d05fff8fa446
*/
const KEY_RIGHT_ALT = 346

/*
C documentation : [GLFW_KEY_RIGHT_SUPER]

[GLFW_KEY_RIGHT_SUPER]: https://www.glfw.org/docs/latest/group__keys.html#gad4547a3e8e247594acb60423fe6502db
*/
const KEY_RIGHT_SUPER = 347

/*
C documentation : [GLFW_KEY_MENU]

[GLFW_KEY_MENU]: https://www.glfw.org/docs/latest/group__keys.html#ga9845be48a745fc232045c9ec174d8820
*/
const KEY_MENU = 348

/*
C documentation : [GLFW_KEY_LAST]

[GLFW_KEY_LAST]: https://www.glfw.org/docs/latest/group__keys.html#ga442cbaef7bfb9a4ba13594dd7fbf2789
*/
const KEY_LAST = KEY_MENU

/*
If this bit is set one or more Shift keys were held down.

C documentation : [GLFW_MOD_SHIFT]

[GLFW_MOD_SHIFT]: https://www.glfw.org/docs/latest/group__mods.html#ga14994d3196c290aaa347248e51740274
*/
const MOD_SHIFT = 0x0001

/*
If this bit is set one or more Control keys were held down.

C documentation : [GLFW_MOD_CONTROL]

[GLFW_MOD_CONTROL]: https://www.glfw.org/docs/latest/group__mods.html#ga6ed94871c3208eefd85713fa929d45aa
*/
const MOD_CONTROL = 0x0002

/*
If this bit is set one or more Alt keys were held down.

C documentation : [GLFW_MOD_ALT]

[GLFW_MOD_ALT]: https://www.glfw.org/docs/latest/group__mods.html#gad2acd5633463c29e07008687ea73c0f4
*/
const MOD_ALT = 0x0004

/*
If this bit is set one or more Super keys were held down.

C documentation : [GLFW_MOD_SUPER]

[GLFW_MOD_SUPER]: https://www.glfw.org/docs/latest/group__mods.html#ga6b64ba10ea0227cf6f42efd0a220aba1
*/
const MOD_SUPER = 0x0008

/*
If this bit is set the Caps Lock key is enabled and the [LOCK_KEY_MODS] input mode is set.

C documentation : [GLFW_MOD_CAPS_LOCK]

[GLFW_MOD_CAPS_LOCK]: https://www.glfw.org/docs/latest/group__mods.html#gaefeef8fcf825a6e43e241b337897200f
*/
const MOD_CAPS_LOCK = 0x0010

/*
If this bit is set the Num Lock key is enabled and the [LOCK_KEY_MODS] input mode is set.

C documentation : [GLFW_MOD_NUM_LOCK]

[GLFW_MOD_NUM_LOCK]: https://www.glfw.org/docs/latest/group__mods.html#ga64e020b8a42af8376e944baf61feecbe
*/
const MOD_NUM_LOCK = 0x0020

/*
See [mouse button input] for how these are used.

C documentation : [GLFW_MOUSE_BUTTON_1]

[mouse button input]: https://www.glfw.org/docs/latest/input_guide.html#input_mouse_button
[GLFW_MOUSE_BUTTON_1]: https://www.glfw.org/docs/latest/group__buttons.html#ga181a6e875251fd8671654eff00f9112e
*/
const MOUSE_BUTTON_1 = 0

/*
C documentation : [GLFW_MOUSE_BUTTON_2]

[GLFW_MOUSE_BUTTON_2]: https://www.glfw.org/docs/latest/group__buttons.html#ga604b39b92c88ce9bd332e97fc3f4156c
*/
const MOUSE_BUTTON_2 = 1

/*
C documentation : [GLFW_MOUSE_BUTTON_3]

[GLFW_MOUSE_BUTTON_3]: https://www.glfw.org/docs/latest/group__buttons.html#ga0130d505563d0236a6f85545f19e1721
*/
const MOUSE_BUTTON_3 = 2

/*
C documentation : [GLFW_MOUSE_BUTTON_4]

[GLFW_MOUSE_BUTTON_4]: https://www.glfw.org/docs/latest/group__buttons.html#ga53f4097bb01d5521c7d9513418c91ca9
*/
const MOUSE_BUTTON_4 = 3

/*
C documentation : [GLFW_MOUSE_BUTTON_5]

[GLFW_MOUSE_BUTTON_5]: https://www.glfw.org/docs/latest/group__buttons.html#gaf08c4ddecb051d3d9667db1d5e417c9c
*/
const MOUSE_BUTTON_5 = 4

/*
C documentation : [GLFW_MOUSE_BUTTON_6]

[GLFW_MOUSE_BUTTON_6]: https://www.glfw.org/docs/latest/group__buttons.html#gae8513e06aab8aa393b595f22c6d8257a
*/
const MOUSE_BUTTON_6 = 5

/*
C documentation : [GLFW_MOUSE_BUTTON_7]

[GLFW_MOUSE_BUTTON_7]: https://www.glfw.org/docs/latest/group__buttons.html#ga8b02a1ab55dde45b3a3883d54ffd7dc7
*/
const MOUSE_BUTTON_7 = 6

/*
C documentation : [GLFW_MOUSE_BUTTON_8]

[GLFW_MOUSE_BUTTON_8]: https://www.glfw.org/docs/latest/group__buttons.html#ga35d5c4263e0dc0d0a4731ca6c562f32c
*/
const MOUSE_BUTTON_8 = 7

/*
C documentation : [GLFW_MOUSE_BUTTON_LAST]

[GLFW_MOUSE_BUTTON_LAST]: https://www.glfw.org/docs/latest/group__buttons.html#gab1fd86a4518a9141ec7bcde2e15a2fdf
*/
const MOUSE_BUTTON_LAST = MOUSE_BUTTON_8

/*
C documentation : [GLFW_MOUSE_BUTTON_LEFT]

[GLFW_MOUSE_BUTTON_LEFT]: https://www.glfw.org/docs/latest/group__buttons.html#gaf37100431dcd5082d48f95ee8bc8cd56
*/
const MOUSE_BUTTON_LEFT = MOUSE_BUTTON_1

/*
C documentation : [GLFW_MOUSE_BUTTON_RIGHT]

[GLFW_MOUSE_BUTTON_RIGHT]: https://www.glfw.org/docs/latest/group__buttons.html#ga3e2f2cf3c4942df73cc094247d275e74
*/
const MOUSE_BUTTON_RIGHT = MOUSE_BUTTON_2

/*
C documentation : [GLFW_MOUSE_BUTTON_MIDDLE]

[GLFW_MOUSE_BUTTON_MIDDLE]: https://www.glfw.org/docs/latest/group__buttons.html#ga34a4d2a701434f763fd93a2ff842b95a
*/
const MOUSE_BUTTON_MIDDLE = MOUSE_BUTTON_3

/*
See [joystick input] for how these are used.

C documentation : [GLFW_JOYSTICK_1]

[joystick input]: https://www.glfw.org/docs/latest/input_guide.html#joystick
[GLFW_JOYSTICK_1]: https://www.glfw.org/docs/latest/group__joysticks.html#ga34a0443d059e9f22272cd4669073f73d
*/
const JOYSTICK_1 = 0

/*
C documentation : [GLFW_JOYSTICK_2]

[GLFW_JOYSTICK_2]: https://www.glfw.org/docs/latest/group__joysticks.html#ga6eab65ec88e65e0850ef8413504cb50c
*/
const JOYSTICK_2 = 1

/*
C documentation : [GLFW_JOYSTICK_3]

[GLFW_JOYSTICK_3]: https://www.glfw.org/docs/latest/group__joysticks.html#gae6f3eedfeb42424c2f5e3161efb0b654
*/
const JOYSTICK_3 = 2

/*
C documentation : [GLFW_JOYSTICK_4]

[GLFW_JOYSTICK_4]: https://www.glfw.org/docs/latest/group__joysticks.html#ga97ddbcad02b7f48d74fad4ddb08fff59
*/
const JOYSTICK_4 = 3

/*
C documentation : [GLFW_JOYSTICK_5]

[GLFW_JOYSTICK_5]: https://www.glfw.org/docs/latest/group__joysticks.html#gae43281bc66d3fa5089fb50c3e7a28695
*/
const JOYSTICK_5 = 4

/*
C documentation : [GLFW_JOYSTICK_6]

[GLFW_JOYSTICK_6]: https://www.glfw.org/docs/latest/group__joysticks.html#ga74771620aa53bd68a487186dea66fd77
*/
const JOYSTICK_6 = 5

/*
C documentation : [GLFW_JOYSTICK_7]

[GLFW_JOYSTICK_7]: https://www.glfw.org/docs/latest/group__joysticks.html#ga20a9f4f3aaefed9ea5e66072fc588b87
*/
const JOYSTICK_7 = 6

/*
C documentation : [GLFW_JOYSTICK_8]

[GLFW_JOYSTICK_8]: https://www.glfw.org/docs/latest/group__joysticks.html#ga21a934c940bcf25db0e4c8fe9b364bdb
*/
const JOYSTICK_8 = 7

/*
C documentation : [GLFW_JOYSTICK_9]

[GLFW_JOYSTICK_9]: https://www.glfw.org/docs/latest/group__joysticks.html#ga87689d47df0ba6f9f5fcbbcaf7b3cecf
*/
const JOYSTICK_9 = 8

/*
C documentation : [GLFW_JOYSTICK_10]

[GLFW_JOYSTICK_10]: https://www.glfw.org/docs/latest/group__joysticks.html#gaef55389ee605d6dfc31aef6fe98c54ec
*/
const JOYSTICK_10 = 9

/*
C documentation : [GLFW_JOYSTICK_11]

[GLFW_JOYSTICK_11]: https://www.glfw.org/docs/latest/group__joysticks.html#gae7d26e3df447c2c14a569fcc18516af4
*/
const JOYSTICK_11 = 10

/*
C documentation : [GLFW_JOYSTICK_12]

[GLFW_JOYSTICK_12]: https://www.glfw.org/docs/latest/group__joysticks.html#gab91bbf5b7ca6be8d3ac5c4d89ff48ac7
*/
const JOYSTICK_12 = 11

/*
C documentation : [GLFW_JOYSTICK_13]

[GLFW_JOYSTICK_13]: https://www.glfw.org/docs/latest/group__joysticks.html#ga5c84fb4e49bf661d7d7c78eb4018c508
*/
const JOYSTICK_13 = 12

/*
C documentation : [GLFW_JOYSTICK_14]

[GLFW_JOYSTICK_14]: https://www.glfw.org/docs/latest/group__joysticks.html#ga89540873278ae5a42b3e70d64164dc74
*/
const JOYSTICK_14 = 13

/*
C documentation : [GLFW_JOYSTICK_15]

[GLFW_JOYSTICK_15]: https://www.glfw.org/docs/latest/group__joysticks.html#ga7b02ab70daf7a78bcc942d5d4cc1dcf9
*/
const JOYSTICK_15 = 14

/*
C documentation : [GLFW_JOYSTICK_16]

[GLFW_JOYSTICK_16]: https://www.glfw.org/docs/latest/group__joysticks.html#ga453edeeabf350827646b6857df4f80ce
*/
const JOYSTICK_16 = 15

/*
C documentation : [GLFW_JOYSTICK_LAST]

[GLFW_JOYSTICK_LAST]: https://www.glfw.org/docs/latest/group__joysticks.html#ga9ca13ebf24c331dd98df17d84a4b72c9
*/
const JOYSTICK_LAST = JOYSTICK_16

/*
See [Gamepad input] for how these are used.

C documentation : [GLFW_GAMEPAD_BUTTON_A]

[Gamepad input]: https://www.glfw.org/docs/latest/input_guide.html#gamepad
[GLFW_GAMEPAD_BUTTON_A]: https://www.glfw.org/docs/latest/group__gamepad__buttons.html#gae055a12fbf4b48b5954c8e1cd129b810
*/
const GAMEPAD_BUTTON_A = 0

/*
C documentation : [GLFW_GAMEPAD_BUTTON_B]

[GLFW_GAMEPAD_BUTTON_B]: https://www.glfw.org/docs/latest/group__gamepad__buttons.html#ga2228a6512fd5950cdb51ba07846546fa
*/
const GAMEPAD_BUTTON_B = 1

/*
C documentation : [GLFW_GAMEPAD_BUTTON_X]

[GLFW_GAMEPAD_BUTTON_X]: https://www.glfw.org/docs/latest/group__gamepad__buttons.html#ga52cc94785cf3fe9a12e246539259887c
*/
const GAMEPAD_BUTTON_X = 2

/*
C documentation : [GLFW_GAMEPAD_BUTTON_Y]

[GLFW_GAMEPAD_BUTTON_Y]: https://www.glfw.org/docs/latest/group__gamepad__buttons.html#gafc931248bda494b530cbe057f386a5ed
*/
const GAMEPAD_BUTTON_Y = 3

/*
C documentation : [GLFW_GAMEPAD_BUTTON_LEFT_BUMPER]

[GLFW_GAMEPAD_BUTTON_LEFT_BUMPER]: https://www.glfw.org/docs/latest/group__gamepad__buttons.html#ga17d67b4f39a39d6b813bd1567a3507c3
*/
const GAMEPAD_BUTTON_LEFT_BUMPER = 4

/*
C documentation : [GLFW_GAMEPAD_BUTTON_RIGHT_BUMPER]

[GLFW_GAMEPAD_BUTTON_RIGHT_BUMPER]: https://www.glfw.org/docs/latest/group__gamepad__buttons.html#gadfbc9ea9bf3aae896b79fa49fdc85c7f
*/
const GAMEPAD_BUTTON_RIGHT_BUMPER = 5

/*
C documentation : [GLFW_GAMEPAD_BUTTON_BACK]

[GLFW_GAMEPAD_BUTTON_BACK]: https://www.glfw.org/docs/latest/group__gamepad__buttons.html#gabc7c0264ce778835b516a472b47f6caf
*/
const GAMEPAD_BUTTON_BACK = 6

/*
C documentation : [GLFW_GAMEPAD_BUTTON_START]

[GLFW_GAMEPAD_BUTTON_START]: https://www.glfw.org/docs/latest/group__gamepad__buttons.html#ga04606949dd9139434b8a1bedf4ac1021
*/
const GAMEPAD_BUTTON_START = 7

/*
C documentation : [GLFW_GAMEPAD_BUTTON_GUIDE]

[GLFW_GAMEPAD_BUTTON_GUIDE]: https://www.glfw.org/docs/latest/group__gamepad__buttons.html#ga7fa48c32e5b2f5db2f080aa0b8b573dc
*/
const GAMEPAD_BUTTON_GUIDE = 8

/*
C documentation : [GLFW_GAMEPAD_BUTTON_LEFT_THUMB]

[GLFW_GAMEPAD_BUTTON_LEFT_THUMB]: https://www.glfw.org/docs/latest/group__gamepad__buttons.html#ga3e089787327454f7bfca7364d6ca206a
*/
const GAMEPAD_BUTTON_LEFT_THUMB = 9

/*
C documentation : [GLFW_GAMEPAD_BUTTON_RIGHT_THUMB]

[GLFW_GAMEPAD_BUTTON_RIGHT_THUMB]: https://www.glfw.org/docs/latest/group__gamepad__buttons.html#ga1c003f52b5aebb45272475b48953b21a
*/
const GAMEPAD_BUTTON_RIGHT_THUMB = 10

/*
C documentation : [GLFW_GAMEPAD_BUTTON_DPAD_UP]

[GLFW_GAMEPAD_BUTTON_DPAD_UP]: https://www.glfw.org/docs/latest/group__gamepad__buttons.html#ga4f1ed6f974a47bc8930d4874a283476a
*/
const GAMEPAD_BUTTON_DPAD_UP = 11

/*
C documentation : [GLFW_GAMEPAD_BUTTON_DPAD_RIGHT]

[GLFW_GAMEPAD_BUTTON_DPAD_RIGHT]: https://www.glfw.org/docs/latest/group__gamepad__buttons.html#gae2a780d2a8c79e0b77c0b7b601ca57c6
*/
const GAMEPAD_BUTTON_DPAD_RIGHT = 12

/*
C documentation : [GLFW_GAMEPAD_BUTTON_DPAD_DOWN]

[GLFW_GAMEPAD_BUTTON_DPAD_DOWN]: https://www.glfw.org/docs/latest/group__gamepad__buttons.html#ga8f2b731b97d80f90f11967a83207665c
*/
const GAMEPAD_BUTTON_DPAD_DOWN = 13

/*
C documentation : [GLFW_GAMEPAD_BUTTON_DPAD_LEFT]

[GLFW_GAMEPAD_BUTTON_DPAD_LEFT]: https://www.glfw.org/docs/latest/group__gamepad__buttons.html#gaf0697e0e8607b2ebe1c93b0c6befe301
*/
const GAMEPAD_BUTTON_DPAD_LEFT = 14

/*
C documentation : [GLFW_GAMEPAD_BUTTON_LAST]

[GLFW_GAMEPAD_BUTTON_LAST]: https://www.glfw.org/docs/latest/group__gamepad__buttons.html#ga5cc98882f4f81dacf761639a567f61eb
*/
const GAMEPAD_BUTTON_LAST = GAMEPAD_BUTTON_DPAD_LEFT

/*
C documentation : [GLFW_GAMEPAD_BUTTON_CROSS]

[GLFW_GAMEPAD_BUTTON_CROSS]: https://www.glfw.org/docs/latest/group__gamepad__buttons.html#gaf08d0df26527c9305253422bd98ed63a
*/
const GAMEPAD_BUTTON_CROSS = GAMEPAD_BUTTON_A

/*
C documentation : [GLFW_GAMEPAD_BUTTON_CIRCLE]

[GLFW_GAMEPAD_BUTTON_CIRCLE]: https://www.glfw.org/docs/latest/group__gamepad__buttons.html#gaaef094b3dacbf15f272b274516839b82
*/
const GAMEPAD_BUTTON_CIRCLE = GAMEPAD_BUTTON_B

/*
C documentation : [GLFW_GAMEPAD_BUTTON_SQUARE]

[GLFW_GAMEPAD_BUTTON_SQUARE]: https://www.glfw.org/docs/latest/group__gamepad__buttons.html#gafc7821e87d77d41ed2cd3e1f726ec35f
*/
const GAMEPAD_BUTTON_SQUARE = GAMEPAD_BUTTON_X

/*
C documentation : [GLFW_GAMEPAD_BUTTON_TRIANGLE]

[GLFW_GAMEPAD_BUTTON_TRIANGLE]: https://www.glfw.org/docs/latest/group__gamepad__buttons.html#ga3a7ef6bcb768a08cd3bf142f7f09f802
*/
const GAMEPAD_BUTTON_TRIANGLE = GAMEPAD_BUTTON_Y

/*
See [Gamepad input] for how these are used.

C documentation : [GLFW_GAMEPAD_AXIS_LEFT_X]

[Gamepad input]: https://www.glfw.org/docs/latest/input_guide.html#gamepad
[GLFW_GAMEPAD_AXIS_LEFT_X]: https://www.glfw.org/docs/latest/group__gamepad__axes.html#ga544e396d092036a7d80c1e5f233f7a38
*/
const GAMEPAD_AXIS_LEFT_X = 0

/*
C documentation : [GLFW_GAMEPAD_AXIS_LEFT_Y]

[GLFW_GAMEPAD_AXIS_LEFT_Y]: https://www.glfw.org/docs/latest/group__gamepad__axes.html#ga64dcf2c6e9be50b7c556ff7671996dd5
*/
const GAMEPAD_AXIS_LEFT_Y = 1

/*
C documentation : [GLFW_GAMEPAD_AXIS_RIGHT_X]

[GLFW_GAMEPAD_AXIS_RIGHT_X]: https://www.glfw.org/docs/latest/group__gamepad__axes.html#gabd6785106cd3c5a044a6e49a395ee2fc
*/
const GAMEPAD_AXIS_RIGHT_X = 2

/*
C documentation : [GLFW_GAMEPAD_AXIS_RIGHT_Y]

[GLFW_GAMEPAD_AXIS_RIGHT_Y]: https://www.glfw.org/docs/latest/group__gamepad__axes.html#ga1cc20566d44d521b7183681a8e88e2e4
*/
const GAMEPAD_AXIS_RIGHT_Y = 3

/*
C documentation : [GLFW_GAMEPAD_AXIS_LEFT_TRIGGER]

[GLFW_GAMEPAD_AXIS_LEFT_TRIGGER]: https://www.glfw.org/docs/latest/group__gamepad__axes.html#ga6d79561dd8907c37354426242901b86e
*/
const GAMEPAD_AXIS_LEFT_TRIGGER = 4

/*
C documentation : [GLFW_GAMEPAD_AXIS_RIGHT_TRIGGER]

[GLFW_GAMEPAD_AXIS_RIGHT_TRIGGER]: https://www.glfw.org/docs/latest/group__gamepad__axes.html#ga121a7d5d20589a423cd1634dd6ee6eab
*/
const GAMEPAD_AXIS_RIGHT_TRIGGER = 5

/*
C documentation : [GLFW_GAMEPAD_AXIS_LAST]

[GLFW_GAMEPAD_AXIS_LAST]: https://www.glfw.org/docs/latest/group__gamepad__axes.html#ga0818fd9433e1359692b7443293e5ac86
*/
const GAMEPAD_AXIS_LAST = GAMEPAD_AXIS_RIGHT_TRIGGER

/*
No error has occurred.

C documentation : [GLFW_NO_ERROR]

# analysis

Yay.

[GLFW_NO_ERROR]: https://www.glfw.org/docs/latest/group__errors.html#gafa30deee5db4d69c4c93d116ed87dbf4
*/
const NO_ERROR = 0

/*
This occurs if a GLFW function was called that must not be called unless the
library is [initialized].

C documentation : [GLFW_NOT_INITIALIZED]

# analysis

Application programmer error.  Initialize GLFW before calling any
function that requires initialization.

[initialized]: https://www.glfw.org/docs/latest/intro_guide.html#intro_init
[GLFW_NOT_INITIALIZED]: https://www.glfw.org/docs/latest/group__errors.html#ga2374ee02c177f12e1fa76ff3ed15e14a
*/
const NOT_INITIALIZED = 0x00010001

/*
This occurs if a GLFW function was called that needs and operates on the
current OpenGL or OpenGL ES context but no context is current on the calling
thread.  One such function is [SwapInterval].

C documentation : [GLFW_NO_CURRENT_CONTEXT]

# analysis

Application programmer error.  Ensure a context is current before
calling functions that require a current context.

[GLFW_NO_CURRENT_CONTEXT]: https://www.glfw.org/docs/latest/group__errors.html#gaa8290386e9528ccb9e42a3a4e16fc0d0
*/
const NO_CURRENT_CONTEXT = 0x00010002

/*
One of the arguments to the function was an invalid enum value, for example
requesting [RED_BITS] with [GetAttrib].

C documentation : [GLFW_INVALID_ENUM]

# analysis

Application programmer error.  Fix the offending call.

[GLFW_INVALID_ENUM]: https://www.glfw.org/docs/latest/group__errors.html#ga76f6bb9c4eea73db675f096b404593ce
*/
const INVALID_ENUM = 0x00010003

/*
One of the arguments to the function was an invalid value, for example
requesting a non-existent OpenGL or OpenGL ES version like 2.7.

Requesting a valid but unavailable OpenGL or OpenGL ES version will instead
result in a [VERSION_UNAVAILABLE] error.

C documentation : [GLFW_INVALID_VALUE]

# analysis

Application programmer error.  Fix the offending call.

[GLFW_INVALID_VALUE]: https://www.glfw.org/docs/latest/group__errors.html#gaaf2ef9aa8202c2b82ac2d921e554c687
*/
const INVALID_VALUE = 0x00010004

/*
A memory allocation failed.

C documentation : [GLFW_OUT_OF_MEMORY]

# analysis

A bug in GLFW or the underlying operating system.  Report the bug
to our [issue tracker].

[issue tracker]: https://github.com/glfw/glfw/issues
[GLFW_OUT_OF_MEMORY]: https://www.glfw.org/docs/latest/group__errors.html#ga9023953a2bcb98c2906afd071d21ee7f
*/
const OUT_OF_MEMORY = 0x00010005

/*
GLFW could not find support for the requested API on the system.

C documentation : [GLFW_API_UNAVAILABLE]

# analysis

The installed graphics driver does not support the requested
API, or does not support it via the chosen context creation API.
Below are a few examples.

Some pre-installed Windows graphics drivers do not support OpenGL.  AMD only
supports OpenGL ES via EGL, while Nvidia and Intel only support it via
a WGL or GLX extension.  macOS does not provide OpenGL ES at all.  The Mesa
EGL, OpenGL and OpenGL ES libraries do not interface with the Nvidia binary
driver.  Older graphics drivers do not support Vulkan.

[GLFW_API_UNAVAILABLE]: https://www.glfw.org/docs/latest/group__errors.html#ga56882b290db23261cc6c053c40c2d08e
*/
const API_UNAVAILABLE = 0x00010006

/*
The requested OpenGL or OpenGL ES version (including any requested context
or framebuffer hints) is not available on this machine.

C documentation : [GLFW_VERSION_UNAVAILABLE]

# analysis

The machine does not support your requirements.  If your
application is sufficiently flexible, downgrade your requirements and try
again.  Otherwise, inform the user that their machine does not match your
requirements.

Future invalid OpenGL and OpenGL ES versions, for example OpenGL 4.8 if 5.0
comes out before the 4.x series gets that far, also fail with this error and
not [INVALID_VALUE], because GLFW cannot know what future versions
will exist.

[GLFW_VERSION_UNAVAILABLE]: https://www.glfw.org/docs/latest/group__errors.html#gad16c5565b4a69f9c2a9ac2c0dbc89462
*/
const VERSION_UNAVAILABLE = 0x00010007

/*
more specific categories.

A platform-specific error occurred that does not match any of the more
specific categories.

C documentation : [GLFW_PLATFORM_ERROR]

# analysis

A bug or configuration error in GLFW, the underlying operating
system or its drivers, or a lack of required resources.  Report the issue to
our [issue tracker].

[issue tracker]: https://github.com/glfw/glfw/issues
[GLFW_PLATFORM_ERROR]: https://www.glfw.org/docs/latest/group__errors.html#gad44162d78100ea5e87cdd38426b8c7a1
*/
const PLATFORM_ERROR = 0x00010008

/*
If emitted during window creation, the requested pixel format is not
supported.

If emitted when querying the clipboard, the contents of the clipboard could
not be converted to the requested format.

C documentation : [GLFW_FORMAT_UNAVAILABLE]

# analysis

If emitted during window creation, one or more
[hard constraints] did not match any of the
available pixel formats.  If your application is sufficiently flexible,
downgrade your requirements and try again.  Otherwise, inform the user that
their machine does not match your requirements.

If emitted when querying the clipboard, ignore the error or report it to
the user, as appropriate.

[hard constraints]: https://www.glfw.org/docs/latest/window_guide.html#window_hints_hard
[GLFW_FORMAT_UNAVAILABLE]: https://www.glfw.org/docs/latest/group__errors.html#ga196e125ef261d94184e2b55c05762f14
*/
const FORMAT_UNAVAILABLE = 0x00010009

/*
A window that does not have an OpenGL or OpenGL ES context was passed to
a function that requires it to have one.

C documentation : [GLFW_NO_WINDOW_CONTEXT]

# analysis

Application programmer error.  Fix the offending call.

[GLFW_NO_WINDOW_CONTEXT]: https://www.glfw.org/docs/latest/group__errors.html#gacff24d2757da752ae4c80bf452356487
*/
const NO_WINDOW_CONTEXT = 0x0001000A

/*
The specified standard cursor shape is not available, either because the
current platform cursor theme does not provide it or because it is not
available on the platform.

C documentation : [GLFW_CURSOR_UNAVAILABLE]

# analysis

Platform or system settings limitation.  Pick another
[standard cursor shape] or create a
[custom cursor].

[standard cursor shape]: https://www.glfw.org/docs/latest/group__shapes.html
[custom cursor]: https://www.glfw.org/docs/latest/input_guide.html#cursor_custom
[GLFW_CURSOR_UNAVAILABLE]: https://www.glfw.org/docs/latest/group__errors.html#ga09d6943923a70ddef3a085f5baee786c
*/
const CURSOR_UNAVAILABLE = 0x0001000B

/*
The requested feature is not provided by the platform, so GLFW is unable to
implement it.  The documentation for each function notes if it could emit
this error.

C documentation : [GLFW_FEATURE_UNAVAILABLE]

# analysis

Platform or platform version limitation.  The error can be ignored
unless the feature is critical to the application.

A function call that emits this error has no effect other than the error and
updating any existing out parameters.

[GLFW_FEATURE_UNAVAILABLE]: https://www.glfw.org/docs/latest/group__errors.html#ga526fba20a01504a8086c763b6ca53ce5
*/
const FEATURE_UNAVAILABLE = 0x0001000C

/*
The requested feature has not yet been implemented in GLFW for this platform.

C documentation : [GLFW_FEATURE_UNIMPLEMENTED]

# analysis

An incomplete implementation of GLFW for this platform, hopefully
fixed in a future release.  The error can be ignored unless the feature is
critical to the application.

A function call that emits this error has no effect other than the error and
updating any existing out parameters.

[GLFW_FEATURE_UNIMPLEMENTED]: https://www.glfw.org/docs/latest/group__errors.html#ga5dda77e023e83151e8bd55a6758f946a
*/
const FEATURE_UNIMPLEMENTED = 0x0001000D

/*
If emitted during initialization, no matching platform was found.  If the [PLATFORM] init hint was set to `GLFW_ANY_PLATFORM`, GLFW could not detect any of
the platforms supported by this library binary, except for the Null platform.  If the
init hint was set to a specific platform, it is either not supported by this library
binary or GLFW was not able to detect it.

If emitted by a native access function, GLFW was initialized for a different platform
than the function is for.

C documentation : [GLFW_PLATFORM_UNAVAILABLE]

# analysis

Failure to detect any platform usually only happens on non-macOS Unix
systems, either when no window system is running or the program was run from
a terminal that does not have the necessary environment variables.  Fall back to
a different platform if possible or notify the user that no usable platform was
detected.

Failure to detect a specific platform may have the same cause as above or be because
support for that platform was not compiled in.  Call [PlatformSupported] to
check whether a specific platform is supported by a library binary.

[GLFW_PLATFORM_UNAVAILABLE]: https://www.glfw.org/docs/latest/group__errors.html#ga3608c6c29ab7a72f3bf019f4c3a2563d
*/
const PLATFORM_UNAVAILABLE = 0x0001000E

/*
Input focus [window hint] or
[window attribute].

C documentation : [GLFW_FOCUSED]

[window hint]: https://www.glfw.org/docs/latest/window_guide.html#GLFW_FOCUSED_hint
[window attribute]: https://www.glfw.org/docs/latest/window_guide.html#GLFW_FOCUSED_attrib
[GLFW_FOCUSED]: https://www.glfw.org/docs/latest/group__window.html#ga54ddb14825a1541a56e22afb5f832a9e
*/
const FOCUSED = 0x00020001

/*
Window iconification [window attribute].

C documentation : [GLFW_ICONIFIED]

[window attribute]: https://www.glfw.org/docs/latest/window_guide.html#GLFW_ICONIFIED_attrib
[GLFW_ICONIFIED]: https://www.glfw.org/docs/latest/group__window.html#ga39d44b7c056e55e581355a92d240b58a
*/
const ICONIFIED = 0x00020002

/*
Window resize-ability [window hint] and
[window attribute].

C documentation : [GLFW_RESIZABLE]

[window hint]: https://www.glfw.org/docs/latest/window_guide.html#GLFW_RESIZABLE_hint
[window attribute]: https://www.glfw.org/docs/latest/window_guide.html#GLFW_RESIZABLE_attrib
[GLFW_RESIZABLE]: https://www.glfw.org/docs/latest/group__window.html#gadba13c7a1b3aa40831eb2beedbd5bd1d
*/
const RESIZABLE = 0x00020003

/*
Window visibility [window hint] and
[window attribute].

C documentation : [GLFW_VISIBLE]

[window hint]: https://www.glfw.org/docs/latest/window_guide.html#GLFW_VISIBLE_hint
[window attribute]: https://www.glfw.org/docs/latest/window_guide.html#GLFW_VISIBLE_attrib
[GLFW_VISIBLE]: https://www.glfw.org/docs/latest/group__window.html#gafb3cdc45297e06d8f1eb13adc69ca6c4
*/
const VISIBLE = 0x00020004

/*
Window decoration [window hint] and
[window attribute].

C documentation : [GLFW_DECORATED]

[window hint]: https://www.glfw.org/docs/latest/window_guide.html#GLFW_DECORATED_hint
[window attribute]: https://www.glfw.org/docs/latest/window_guide.html#GLFW_DECORATED_attrib
[GLFW_DECORATED]: https://www.glfw.org/docs/latest/group__window.html#ga21b854d36314c94d65aed84405b2f25e
*/
const DECORATED = 0x00020005

/*
Window auto-iconification [window hint] and
[window attribute].

C documentation : [GLFW_AUTO_ICONIFY]

[window hint]: https://www.glfw.org/docs/latest/window_guide.html#GLFW_AUTO_ICONIFY_hint
[window attribute]: https://www.glfw.org/docs/latest/window_guide.html#GLFW_AUTO_ICONIFY_attrib
[GLFW_AUTO_ICONIFY]: https://www.glfw.org/docs/latest/group__window.html#ga9d9874fc928200136a6dcdad726aa252
*/
const AUTO_ICONIFY = 0x00020006

/*
Window decoration [window hint] and
[window attribute].

C documentation : [GLFW_FLOATING]

[window hint]: https://www.glfw.org/docs/latest/window_guide.html#GLFW_FLOATING_hint
[window attribute]: https://www.glfw.org/docs/latest/window_guide.html#GLFW_FLOATING_attrib
[GLFW_FLOATING]: https://www.glfw.org/docs/latest/group__window.html#ga7fb0be51407783b41adbf5bec0b09d80
*/
const FLOATING = 0x00020007

/*
Window maximization [window hint] and
[window attribute].

C documentation : [GLFW_MAXIMIZED]

[window hint]: https://www.glfw.org/docs/latest/window_guide.html#GLFW_MAXIMIZED_hint
[window attribute]: https://www.glfw.org/docs/latest/window_guide.html#GLFW_MAXIMIZED_attrib
[GLFW_MAXIMIZED]: https://www.glfw.org/docs/latest/group__window.html#gad8ccb396253ad0b72c6d4c917eb38a03
*/
const MAXIMIZED = 0x00020008

/*
Cursor centering [window hint].

C documentation : [GLFW_CENTER_CURSOR]

[window hint]: https://www.glfw.org/docs/latest/window_guide.html#GLFW_CENTER_CURSOR_hint
[GLFW_CENTER_CURSOR]: https://www.glfw.org/docs/latest/group__window.html#ga5ac0847c0aa0b3619f2855707b8a7a77
*/
const CENTER_CURSOR = 0x00020009

/*
Window framebuffer transparency
[window hint] and
[window attribute].

C documentation : [GLFW_TRANSPARENT_FRAMEBUFFER]

[window hint]: https://www.glfw.org/docs/latest/window_guide.html#GLFW_TRANSPARENT_FRAMEBUFFER_hint
[window attribute]: https://www.glfw.org/docs/latest/window_guide.html#GLFW_TRANSPARENT_FRAMEBUFFER_attrib
[GLFW_TRANSPARENT_FRAMEBUFFER]: https://www.glfw.org/docs/latest/group__window.html#ga60a0578c3b9449027d683a9c6abb9f14
*/
const TRANSPARENT_FRAMEBUFFER = 0x0002000A

/*
Mouse cursor hover [window attribute].

C documentation : [GLFW_HOVERED]

[window attribute]: https://www.glfw.org/docs/latest/window_guide.html#GLFW_HOVERED_attrib
[GLFW_HOVERED]: https://www.glfw.org/docs/latest/group__window.html#ga8665c71c6fa3d22425c6a0e8a3f89d8a
*/
const HOVERED = 0x0002000B

/*
Input focus [window hint] or
[window attribute].

C documentation : [GLFW_FOCUS_ON_SHOW]

[window hint]: https://www.glfw.org/docs/latest/window_guide.html#GLFW_FOCUS_ON_SHOW_hint
[window attribute]: https://www.glfw.org/docs/latest/window_guide.html#GLFW_FOCUS_ON_SHOW_attrib
[GLFW_FOCUS_ON_SHOW]: https://www.glfw.org/docs/latest/group__window.html#gafa94b1da34bfd6488c0d709761504dfc
*/
const FOCUS_ON_SHOW = 0x0002000C

/*
Mouse input transparency [window hint] or
[window attribute].

C documentation : [GLFW_MOUSE_PASSTHROUGH]

[window hint]: https://www.glfw.org/docs/latest/window_guide.html#GLFW_MOUSE_PASSTHROUGH_hint
[window attribute]: https://www.glfw.org/docs/latest/window_guide.html#GLFW_MOUSE_PASSTHROUGH_attrib
[GLFW_MOUSE_PASSTHROUGH]: https://www.glfw.org/docs/latest/group__window.html#ga88981797d29800808ec242274ab5c03a
*/
const MOUSE_PASSTHROUGH = 0x0002000D

/*
Initial position x-coordinate [window hint].

C documentation : [GLFW_POSITION_X]

[window hint]: https://www.glfw.org/docs/latest/group__window.html#gaededa6b208b8e31343da56bb349c6fb2
[GLFW_POSITION_X]: https://www.glfw.org/docs/latest/group__window.html#gaededa6b208b8e31343da56bb349c6fb2
*/
const POSITION_X = 0x0002000E

/*
Initial position y-coordinate [window hint].

C documentation : [GLFW_POSITION_Y]

[window hint]: https://www.glfw.org/docs/latest/group__window.html#ga6b3ccf63683c81f479e2a98f5027200e
[GLFW_POSITION_Y]: https://www.glfw.org/docs/latest/group__window.html#ga6b3ccf63683c81f479e2a98f5027200e
*/
const POSITION_Y = 0x0002000F

/*
Framebuffer bit depth [hint].

C documentation : [GLFW_RED_BITS]

[hint]: https://www.glfw.org/docs/latest/group__window.html#gaf78ed8e417dbcc1e354906cc2708c982
[GLFW_RED_BITS]: https://www.glfw.org/docs/latest/group__window.html#gaf78ed8e417dbcc1e354906cc2708c982
*/
const RED_BITS = 0x00021001

/*
Framebuffer bit depth [hint].

C documentation : [GLFW_GREEN_BITS]

[hint]: https://www.glfw.org/docs/latest/group__window.html#gafba3b72638c914e5fb8a237dd4c50d4d
[GLFW_GREEN_BITS]: https://www.glfw.org/docs/latest/group__window.html#gafba3b72638c914e5fb8a237dd4c50d4d
*/
const GREEN_BITS = 0x00021002

/*
Framebuffer bit depth [hint].

C documentation : [GLFW_BLUE_BITS]

[hint]: https://www.glfw.org/docs/latest/group__window.html#gab292ea403db6d514537b515311bf9ae3
[GLFW_BLUE_BITS]: https://www.glfw.org/docs/latest/group__window.html#gab292ea403db6d514537b515311bf9ae3
*/
const BLUE_BITS = 0x00021003

/*
Framebuffer bit depth [hint].

C documentation : [GLFW_ALPHA_BITS]

[hint]: https://www.glfw.org/docs/latest/group__window.html#gafed79a3f468997877da86c449bd43e8c
[GLFW_ALPHA_BITS]: https://www.glfw.org/docs/latest/group__window.html#gafed79a3f468997877da86c449bd43e8c
*/
const ALPHA_BITS = 0x00021004

/*
Framebuffer bit depth [hint].

C documentation : [GLFW_DEPTH_BITS]

[hint]: https://www.glfw.org/docs/latest/group__window.html#ga318a55eac1fee57dfe593b6d38149d07
[GLFW_DEPTH_BITS]: https://www.glfw.org/docs/latest/group__window.html#ga318a55eac1fee57dfe593b6d38149d07
*/
const DEPTH_BITS = 0x00021005

/*
Framebuffer bit depth [hint].

C documentation : [GLFW_STENCIL_BITS]

[hint]: https://www.glfw.org/docs/latest/group__window.html#ga5339890a45a1fb38e93cb9fcc5fd069d
[GLFW_STENCIL_BITS]: https://www.glfw.org/docs/latest/group__window.html#ga5339890a45a1fb38e93cb9fcc5fd069d
*/
const STENCIL_BITS = 0x00021006

/*
Framebuffer bit depth [hint].

C documentation : [GLFW_ACCUM_RED_BITS]

[hint]: https://www.glfw.org/docs/latest/group__window.html#gaead34a9a683b2bc20eecf30ba738bfc6
[GLFW_ACCUM_RED_BITS]: https://www.glfw.org/docs/latest/group__window.html#gaead34a9a683b2bc20eecf30ba738bfc6
*/
const ACCUM_RED_BITS = 0x00021007

/*
Framebuffer bit depth [hint].

C documentation : [GLFW_ACCUM_GREEN_BITS]

[hint]: https://www.glfw.org/docs/latest/group__window.html#ga65713cee1326f8e9d806fdf93187b471
[GLFW_ACCUM_GREEN_BITS]: https://www.glfw.org/docs/latest/group__window.html#ga65713cee1326f8e9d806fdf93187b471
*/
const ACCUM_GREEN_BITS = 0x00021008

/*
Framebuffer bit depth [hint].

C documentation : [GLFW_ACCUM_BLUE_BITS]

[hint]: https://www.glfw.org/docs/latest/group__window.html#ga22bbe9104a8ce1f8b88fb4f186aa36ce
[GLFW_ACCUM_BLUE_BITS]: https://www.glfw.org/docs/latest/group__window.html#ga22bbe9104a8ce1f8b88fb4f186aa36ce
*/
const ACCUM_BLUE_BITS = 0x00021009

/*
Framebuffer bit depth [hint].

C documentation : [GLFW_ACCUM_ALPHA_BITS]

[hint]: https://www.glfw.org/docs/latest/group__window.html#gae829b55591c18169a40ab4067a041b1f
[GLFW_ACCUM_ALPHA_BITS]: https://www.glfw.org/docs/latest/group__window.html#gae829b55591c18169a40ab4067a041b1f
*/
const ACCUM_ALPHA_BITS = 0x0002100A

/*
Framebuffer auxiliary buffer [hint].

C documentation : [GLFW_AUX_BUFFERS]

[hint]: https://www.glfw.org/docs/latest/group__window.html#gab05108c5029443b371112b031d1fa174
[GLFW_AUX_BUFFERS]: https://www.glfw.org/docs/latest/group__window.html#gab05108c5029443b371112b031d1fa174
*/
const AUX_BUFFERS = 0x0002100B

/*
OpenGL stereoscopic rendering [hint].

C documentation : [GLFW_STEREO]

[hint]: https://www.glfw.org/docs/latest/group__window.html#ga83d991efca02537e2d69969135b77b03
[GLFW_STEREO]: https://www.glfw.org/docs/latest/group__window.html#ga83d991efca02537e2d69969135b77b03
*/
const STEREO = 0x0002100C

/*
Framebuffer MSAA samples [hint].

C documentation : [GLFW_SAMPLES]

[hint]: https://www.glfw.org/docs/latest/group__window.html#ga2cdf86fdcb7722fb8829c4e201607535
[GLFW_SAMPLES]: https://www.glfw.org/docs/latest/group__window.html#ga2cdf86fdcb7722fb8829c4e201607535
*/
const SAMPLES = 0x0002100D

/*
Framebuffer sRGB [hint].

C documentation : [GLFW_SRGB_CAPABLE]

[hint]: https://www.glfw.org/docs/latest/group__window.html#ga444a8f00414a63220591f9fdb7b5642b
[GLFW_SRGB_CAPABLE]: https://www.glfw.org/docs/latest/group__window.html#ga444a8f00414a63220591f9fdb7b5642b
*/
const SRGB_CAPABLE = 0x0002100E

/*
Monitor refresh rate [hint].

C documentation : [GLFW_REFRESH_RATE]

[hint]: https://www.glfw.org/docs/latest/group__window.html#ga0f20825e6e47ee8ba389024519682212
[GLFW_REFRESH_RATE]: https://www.glfw.org/docs/latest/group__window.html#ga0f20825e6e47ee8ba389024519682212
*/
const REFRESH_RATE = 0x0002100F

/*
Framebuffer double buffering [hint] and
[attribute].

C documentation : [GLFW_DOUBLEBUFFER]

[hint]: https://www.glfw.org/docs/latest/window_guide.html#GLFW_DOUBLEBUFFER_hint
[attribute]: https://www.glfw.org/docs/latest/window_guide.html#GLFW_DOUBLEBUFFER_attrib
[GLFW_DOUBLEBUFFER]: https://www.glfw.org/docs/latest/group__window.html#ga714a5d569e8a274ea58fdfa020955339
*/
const DOUBLEBUFFER = 0x00021010

/*
Context client API [hint] and
[attribute].

C documentation : [GLFW_CLIENT_API]

[hint]: https://www.glfw.org/docs/latest/window_guide.html#GLFW_CLIENT_API_hint
[attribute]: https://www.glfw.org/docs/latest/window_guide.html#GLFW_CLIENT_API_attrib
[GLFW_CLIENT_API]: https://www.glfw.org/docs/latest/group__window.html#ga649309cf72a3d3de5b1348ca7936c95b
*/
const CLIENT_API = 0x00022001

/*
Context client API major version [hint]
and [attribute].

C documentation : [GLFW_CONTEXT_VERSION_MAJOR]

[hint]: https://www.glfw.org/docs/latest/window_guide.html#GLFW_CONTEXT_VERSION_MAJOR_hint
[attribute]: https://www.glfw.org/docs/latest/window_guide.html#GLFW_CONTEXT_VERSION_MAJOR_attrib
[GLFW_CONTEXT_VERSION_MAJOR]: https://www.glfw.org/docs/latest/group__window.html#gafe5e4922de1f9932d7e9849bb053b0c0
*/
const CONTEXT_VERSION_MAJOR = 0x00022002

/*
Context client API minor version [hint]
and [attribute].

C documentation : [GLFW_CONTEXT_VERSION_MINOR]

[hint]: https://www.glfw.org/docs/latest/window_guide.html#GLFW_CONTEXT_VERSION_MINOR_hint
[attribute]: https://www.glfw.org/docs/latest/window_guide.html#GLFW_CONTEXT_VERSION_MINOR_attrib
[GLFW_CONTEXT_VERSION_MINOR]: https://www.glfw.org/docs/latest/group__window.html#ga31aca791e4b538c4e4a771eb95cc2d07
*/
const CONTEXT_VERSION_MINOR = 0x00022003

/*
Context client API revision number
[attribute].

C documentation : [GLFW_CONTEXT_REVISION]

[attribute]: https://www.glfw.org/docs/latest/window_guide.html#GLFW_CONTEXT_REVISION_attrib
[GLFW_CONTEXT_REVISION]: https://www.glfw.org/docs/latest/group__window.html#gafb9475071aa77c6fb05ca5a5c8678a08
*/
const CONTEXT_REVISION = 0x00022004

/*
Context client API revision number [hint]
and [attribute].

C documentation : [GLFW_CONTEXT_ROBUSTNESS]

[hint]: https://www.glfw.org/docs/latest/window_guide.html#GLFW_CONTEXT_ROBUSTNESS_hint
[attribute]: https://www.glfw.org/docs/latest/window_guide.html#GLFW_CONTEXT_ROBUSTNESS_attrib
[GLFW_CONTEXT_ROBUSTNESS]: https://www.glfw.org/docs/latest/group__window.html#gade3593916b4c507900aa2d6844810e00
*/
const CONTEXT_ROBUSTNESS = 0x00022005

/*
OpenGL forward-compatibility [hint]
and [attribute].

C documentation : [GLFW_OPENGL_FORWARD_COMPAT]

[hint]: https://www.glfw.org/docs/latest/window_guide.html#GLFW_OPENGL_FORWARD_COMPAT_hint
[attribute]: https://www.glfw.org/docs/latest/window_guide.html#GLFW_OPENGL_FORWARD_COMPAT_attrib
[GLFW_OPENGL_FORWARD_COMPAT]: https://www.glfw.org/docs/latest/group__window.html#ga13d24b12465da8b28985f46c8557925b
*/
const OPENGL_FORWARD_COMPAT = 0x00022006

/*
Debug mode context [hint] and
[attribute].

C documentation : [GLFW_CONTEXT_DEBUG]

[hint]: https://www.glfw.org/docs/latest/window_guide.html#GLFW_CONTEXT_DEBUG_hint
[attribute]: https://www.glfw.org/docs/latest/window_guide.html#GLFW_CONTEXT_DEBUG_attrib
[GLFW_CONTEXT_DEBUG]: https://www.glfw.org/docs/latest/group__window.html#ga8d55e3afec73c7de0509c3b7ad1d9e3f
*/
const CONTEXT_DEBUG = 0x00022007

/*
This is an alias for compatibility with earlier versions.

C documentation : [GLFW_OPENGL_DEBUG_CONTEXT]

[GLFW_OPENGL_DEBUG_CONTEXT]: https://www.glfw.org/docs/latest/group__window.html#ga87ec2df0b915201e950ca42d5d0831e1
*/
const OPENGL_DEBUG_CONTEXT = CONTEXT_DEBUG

/*
OpenGL profile [hint] and
[attribute].

C documentation : [GLFW_OPENGL_PROFILE]

[hint]: https://www.glfw.org/docs/latest/window_guide.html#GLFW_OPENGL_PROFILE_hint
[attribute]: https://www.glfw.org/docs/latest/window_guide.html#GLFW_OPENGL_PROFILE_attrib
[GLFW_OPENGL_PROFILE]: https://www.glfw.org/docs/latest/group__window.html#ga44f3a6b4261fbe351e0b950b0f372e12
*/
const OPENGL_PROFILE = 0x00022008

/*
Context flush-on-release [hint] and
[attribute].

C documentation : [GLFW_CONTEXT_RELEASE_BEHAVIOR]

[hint]: https://www.glfw.org/docs/latest/window_guide.html#GLFW_CONTEXT_RELEASE_BEHAVIOR_hint
[attribute]: https://www.glfw.org/docs/latest/window_guide.html#GLFW_CONTEXT_RELEASE_BEHAVIOR_attrib
[GLFW_CONTEXT_RELEASE_BEHAVIOR]: https://www.glfw.org/docs/latest/group__window.html#ga72b648a8378fe3310c7c7bbecc0f7be6
*/
const CONTEXT_RELEASE_BEHAVIOR = 0x00022009

/*
Context error suppression [hint] and
[attribute].

C documentation : [GLFW_CONTEXT_NO_ERROR]

[hint]: https://www.glfw.org/docs/latest/window_guide.html#GLFW_CONTEXT_NO_ERROR_hint
[attribute]: https://www.glfw.org/docs/latest/window_guide.html#GLFW_CONTEXT_NO_ERROR_attrib
[GLFW_CONTEXT_NO_ERROR]: https://www.glfw.org/docs/latest/group__window.html#ga5a52fdfd46d8249c211f923675728082
*/
const CONTEXT_NO_ERROR = 0x0002200A

/*
Context creation API [hint] and
[attribute].

C documentation : [GLFW_CONTEXT_CREATION_API]

[hint]: https://www.glfw.org/docs/latest/window_guide.html#GLFW_CONTEXT_CREATION_API_hint
[attribute]: https://www.glfw.org/docs/latest/window_guide.html#GLFW_CONTEXT_CREATION_API_attrib
[GLFW_CONTEXT_CREATION_API]: https://www.glfw.org/docs/latest/group__window.html#ga5154cebfcd831c1cc63a4d5ac9bb4486
*/
const CONTEXT_CREATION_API = 0x0002200B

/*
[window hint].

C documentation : [GLFW_SCALE_TO_MONITOR]

[window hint]: https://www.glfw.org/docs/latest/group__window.html#ga620bc4280c7eab81ac9f02204500ed47
[GLFW_SCALE_TO_MONITOR]: https://www.glfw.org/docs/latest/group__window.html#ga620bc4280c7eab81ac9f02204500ed47
*/
const SCALE_TO_MONITOR = 0x0002200C

/*
[window hint].

C documentation : [GLFW_SCALE_FRAMEBUFFER]

[window hint]: https://www.glfw.org/docs/latest/window_guide.html#GLFW_SCALE_FRAMEBUFFER_hint
[GLFW_SCALE_FRAMEBUFFER]: https://www.glfw.org/docs/latest/group__window.html#gaa5a9c6b4722670fd33d6e8a88f2e21bc
*/
const SCALE_FRAMEBUFFER = 0x0002200D

/*
This is an alias for the
[GLFW_SCALE_FRAMEBUFFER] window hint for
compatibility with earlier versions.

C documentation : [GLFW_COCOA_RETINA_FRAMEBUFFER]

[GLFW_SCALE_FRAMEBUFFER]: https://www.glfw.org/docs/latest/window_guide.html#GLFW_SCALE_FRAMEBUFFER_hint
[GLFW_COCOA_RETINA_FRAMEBUFFER]: https://www.glfw.org/docs/latest/group__window.html#gab6ef2d02eb55800d249ccf1af253c35e
*/
const COCOA_RETINA_FRAMEBUFFER = 0x00023001

/*
[window hint].

C documentation : [GLFW_COCOA_FRAME_NAME]

[window hint]: https://www.glfw.org/docs/latest/window_guide.html#GLFW_COCOA_FRAME_NAME_hint
[GLFW_COCOA_FRAME_NAME]: https://www.glfw.org/docs/latest/group__window.html#ga70fa0fbc745de6aa824df79a580e84b5
*/
const COCOA_FRAME_NAME = 0x00023002

/*
[window hint].

C documentation : [GLFW_COCOA_GRAPHICS_SWITCHING]

[window hint]: https://www.glfw.org/docs/latest/window_guide.html#GLFW_COCOA_GRAPHICS_SWITCHING_hint
[GLFW_COCOA_GRAPHICS_SWITCHING]: https://www.glfw.org/docs/latest/group__window.html#ga53c84ed2ddd94e15bbd44b1f6f7feafc
*/
const COCOA_GRAPHICS_SWITCHING = 0x00023003

/*
[window hint].

C documentation : [GLFW_X11_CLASS_NAME]

[window hint]: https://www.glfw.org/docs/latest/window_guide.html#GLFW_X11_CLASS_NAME_hint
[GLFW_X11_CLASS_NAME]: https://www.glfw.org/docs/latest/group__window.html#gae5a9ea2fccccd92edbd343fc56461114
*/
const X11_CLASS_NAME = 0x00024001

/*
[window hint].

C documentation : [GLFW_X11_INSTANCE_NAME]

[window hint]: https://www.glfw.org/docs/latest/window_guide.html#GLFW_X11_CLASS_NAME_hint
[GLFW_X11_INSTANCE_NAME]: https://www.glfw.org/docs/latest/group__window.html#ga494c3c0d911e4b860b946530a3e389e8
*/
const X11_INSTANCE_NAME = 0x00024002

/*
C documentation : [GLFW_WIN32_KEYBOARD_MENU]

[GLFW_WIN32_KEYBOARD_MENU]: https://www.glfw.org/docs/latest/group__window.html#gaf65ea8dafdc0edb07b821b9a336d5043
*/
const WIN32_KEYBOARD_MENU = 0x00025001

/*
C documentation : [GLFW_WIN32_SHOWDEFAULT]

[GLFW_WIN32_SHOWDEFAULT]: https://www.glfw.org/docs/latest/group__window.html#gace10f3846571de62243b46f75d978487
*/
const WIN32_SHOWDEFAULT = 0x00025002

/*
[window hint].

Allows specification of the Wayland app_id.

C documentation : [GLFW_WAYLAND_APP_ID]

[window hint]: https://www.glfw.org/docs/latest/window_guide.html#GLFW_WAYLAND_APP_ID_hint
[GLFW_WAYLAND_APP_ID]: https://www.glfw.org/docs/latest/group__window.html#gafbf1ce7a4362c75e602a4df9e1bdecd3
*/
const WAYLAND_APP_ID = 0x00026001

/*
C documentation : [GLFW_NO_API]

[GLFW_NO_API]: https://www.glfw.org/docs/latest/glfw3_8h.html#a8f6dcdc968d214ff14779564f1389264
*/
const NO_API = 0

/*
C documentation : [GLFW_OPENGL_API]

[GLFW_OPENGL_API]: https://www.glfw.org/docs/latest/glfw3_8h.html#a01b3f66db266341425e9abee6b257db2
*/
const OPENGL_API = 0x00030001

/*
C documentation : [GLFW_OPENGL_ES_API]

[GLFW_OPENGL_ES_API]: https://www.glfw.org/docs/latest/glfw3_8h.html#a28d9b3bc6c2a522d815c8e146595051f
*/
const OPENGL_ES_API = 0x00030002

/*
C documentation : [GLFW_NO_ROBUSTNESS]

[GLFW_NO_ROBUSTNESS]: https://www.glfw.org/docs/latest/glfw3_8h.html#a8b306cb27f5bb0d6d67c7356a0e0fc34
*/
const NO_ROBUSTNESS = 0

/*
C documentation : [GLFW_NO_RESET_NOTIFICATION]

[GLFW_NO_RESET_NOTIFICATION]: https://www.glfw.org/docs/latest/glfw3_8h.html#aee84a679230d205005e22487ff678a85
*/
const NO_RESET_NOTIFICATION = 0x00031001

/*
C documentation : [GLFW_LOSE_CONTEXT_ON_RESET]

[GLFW_LOSE_CONTEXT_ON_RESET]: https://www.glfw.org/docs/latest/glfw3_8h.html#aec1132f245143fc915b2f0995228564c
*/
const LOSE_CONTEXT_ON_RESET = 0x00031002

/*
C documentation : [GLFW_OPENGL_ANY_PROFILE]

[GLFW_OPENGL_ANY_PROFILE]: https://www.glfw.org/docs/latest/glfw3_8h.html#ad6f2335d6f21cc9bab96633b1c111d5f
*/
const OPENGL_ANY_PROFILE = 0

/*
C documentation : [GLFW_OPENGL_CORE_PROFILE]

[GLFW_OPENGL_CORE_PROFILE]: https://www.glfw.org/docs/latest/glfw3_8h.html#af094bb16da76f66ebceb19ee213b3de8
*/
const OPENGL_CORE_PROFILE = 0x00032001

/*
C documentation : [GLFW_OPENGL_COMPAT_PROFILE]

[GLFW_OPENGL_COMPAT_PROFILE]: https://www.glfw.org/docs/latest/glfw3_8h.html#ac06b663d79c8fcf04669cc8fcc0b7670
*/
const OPENGL_COMPAT_PROFILE = 0x00032002

/*
C documentation : [GLFW_CURSOR]

[GLFW_CURSOR]: https://www.glfw.org/docs/latest/glfw3_8h.html#aade31da5b884a84a7625c6b059b9132c
*/
const CURSOR = 0x00033001

/*
C documentation : [GLFW_STICKY_KEYS]

[GLFW_STICKY_KEYS]: https://www.glfw.org/docs/latest/glfw3_8h.html#ae3bbe2315b7691ab088159eb6c9110fc
*/
const STICKY_KEYS = 0x00033002

/*
C documentation : [GLFW_STICKY_MOUSE_BUTTONS]

[GLFW_STICKY_MOUSE_BUTTONS]: https://www.glfw.org/docs/latest/glfw3_8h.html#a4d7ce8ce71030c3b04e2b78145bc59d1
*/
const STICKY_MOUSE_BUTTONS = 0x00033003

/*
C documentation : [GLFW_LOCK_KEY_MODS]

[GLFW_LOCK_KEY_MODS]: https://www.glfw.org/docs/latest/glfw3_8h.html#a07b84de0b52143e1958f88a7d9105947
*/
const LOCK_KEY_MODS = 0x00033004

/*
C documentation : [GLFW_RAW_MOUSE_MOTION]

[GLFW_RAW_MOUSE_MOTION]: https://www.glfw.org/docs/latest/glfw3_8h.html#aeeda1be76a44a1fc97c1282e06281fbb
*/
const RAW_MOUSE_MOTION = 0x00033005

/*
C documentation : [GLFW_CURSOR_NORMAL]

[GLFW_CURSOR_NORMAL]: https://www.glfw.org/docs/latest/glfw3_8h.html#ae04dd25c8577e19fa8c97368561f6c68
*/
const CURSOR_NORMAL = 0x00034001

/*
C documentation : [GLFW_CURSOR_HIDDEN]

[GLFW_CURSOR_HIDDEN]: https://www.glfw.org/docs/latest/glfw3_8h.html#ac4d5cb9d78de8573349c58763d53bf11
*/
const CURSOR_HIDDEN = 0x00034002

/*
C documentation : [GLFW_CURSOR_DISABLED]

[GLFW_CURSOR_DISABLED]: https://www.glfw.org/docs/latest/glfw3_8h.html#a2315b99a329ce53e6a13a9d46fd5ca88
*/
const CURSOR_DISABLED = 0x00034003

/*
C documentation : [GLFW_CURSOR_CAPTURED]

[GLFW_CURSOR_CAPTURED]: https://www.glfw.org/docs/latest/glfw3_8h.html#ac1dbfa0cb4641a0edc93412ade0895dc
*/
const CURSOR_CAPTURED = 0x00034004

/*
C documentation : [GLFW_ANY_RELEASE_BEHAVIOR]

[GLFW_ANY_RELEASE_BEHAVIOR]: https://www.glfw.org/docs/latest/glfw3_8h.html#a6b47d806f285efe9bfd7aeec667297ee
*/
const ANY_RELEASE_BEHAVIOR = 0

/*
C documentation : [GLFW_RELEASE_BEHAVIOR_FLUSH]

[GLFW_RELEASE_BEHAVIOR_FLUSH]: https://www.glfw.org/docs/latest/glfw3_8h.html#a999961d391db49cb4f949c1dece0e13b
*/
const RELEASE_BEHAVIOR_FLUSH = 0x00035001

/*
C documentation : [GLFW_RELEASE_BEHAVIOR_NONE]

[GLFW_RELEASE_BEHAVIOR_NONE]: https://www.glfw.org/docs/latest/glfw3_8h.html#afca09088eccacdce4b59036cfae349c5
*/
const RELEASE_BEHAVIOR_NONE = 0x00035002

/*
C documentation : [GLFW_NATIVE_CONTEXT_API]

[GLFW_NATIVE_CONTEXT_API]: https://www.glfw.org/docs/latest/glfw3_8h.html#a0494c9bfd3f584ab41e6dbeeaa0e6a19
*/
const NATIVE_CONTEXT_API = 0x00036001

/*
C documentation : [GLFW_EGL_CONTEXT_API]

[GLFW_EGL_CONTEXT_API]: https://www.glfw.org/docs/latest/glfw3_8h.html#a03cf65c9ab01fc8b872ba58842c531c9
*/
const EGL_CONTEXT_API = 0x00036002

/*
C documentation : [GLFW_OSMESA_CONTEXT_API]

[GLFW_OSMESA_CONTEXT_API]: https://www.glfw.org/docs/latest/glfw3_8h.html#afd34a473af9fa81f317910ea371b19e3
*/
const OSMESA_CONTEXT_API = 0x00036003

/*
C documentation : [GLFW_ANGLE_PLATFORM_TYPE_NONE]

[GLFW_ANGLE_PLATFORM_TYPE_NONE]: https://www.glfw.org/docs/latest/glfw3_8h.html#ae78e673449c2a2b8c560ca1b1e283228
*/
const ANGLE_PLATFORM_TYPE_NONE = 0x00037001

/*
C documentation : [GLFW_ANGLE_PLATFORM_TYPE_OPENGL]

[GLFW_ANGLE_PLATFORM_TYPE_OPENGL]: https://www.glfw.org/docs/latest/glfw3_8h.html#ad8d9e97ed7790811470366b338833623
*/
const ANGLE_PLATFORM_TYPE_OPENGL = 0x00037002

/*
C documentation : [GLFW_ANGLE_PLATFORM_TYPE_OPENGLES]

[GLFW_ANGLE_PLATFORM_TYPE_OPENGLES]: https://www.glfw.org/docs/latest/glfw3_8h.html#a0003c089da020cbf957218e70245bb65
*/
const ANGLE_PLATFORM_TYPE_OPENGLES = 0x00037003

/*
C documentation : [GLFW_ANGLE_PLATFORM_TYPE_D3D9]

[GLFW_ANGLE_PLATFORM_TYPE_D3D9]: https://www.glfw.org/docs/latest/glfw3_8h.html#a6e8fdc83113d247ad792bb5c4e82c894
*/
const ANGLE_PLATFORM_TYPE_D3D9 = 0x00037004

/*
C documentation : [GLFW_ANGLE_PLATFORM_TYPE_D3D11]

[GLFW_ANGLE_PLATFORM_TYPE_D3D11]: https://www.glfw.org/docs/latest/glfw3_8h.html#ad6eae659811a52a5cdc43c362aedfa33
*/
const ANGLE_PLATFORM_TYPE_D3D11 = 0x00037005

/*
C documentation : [GLFW_ANGLE_PLATFORM_TYPE_VULKAN]

[GLFW_ANGLE_PLATFORM_TYPE_VULKAN]: https://www.glfw.org/docs/latest/glfw3_8h.html#a579ac83506c7546709dad91960cc7ca1
*/
const ANGLE_PLATFORM_TYPE_VULKAN = 0x00037007

/*
C documentation : [GLFW_ANGLE_PLATFORM_TYPE_METAL]

[GLFW_ANGLE_PLATFORM_TYPE_METAL]: https://www.glfw.org/docs/latest/glfw3_8h.html#ab56d91b26cf223dc67590a93a2f8507d
*/
const ANGLE_PLATFORM_TYPE_METAL = 0x00037008

/*
C documentation : [GLFW_WAYLAND_PREFER_LIBDECOR]

[GLFW_WAYLAND_PREFER_LIBDECOR]: https://www.glfw.org/docs/latest/glfw3_8h.html#a92b0d7e0eaeeefaccc0ccc2ccb130e99
*/
const WAYLAND_PREFER_LIBDECOR = 0x00038001

/*
C documentation : [GLFW_WAYLAND_DISABLE_LIBDECOR]

[GLFW_WAYLAND_DISABLE_LIBDECOR]: https://www.glfw.org/docs/latest/glfw3_8h.html#aadcea7c6afbf86b848404457c4253fd7
*/
const WAYLAND_DISABLE_LIBDECOR = 0x00038002

/*
C documentation : [GLFW_ANY_POSITION]

[GLFW_ANY_POSITION]: https://www.glfw.org/docs/latest/glfw3_8h.html#aa0e681bf859ef1bb8355692a70b0ee92
*/
const ANY_POSITION = 0x80000000

/*
The regular arrow cursor shape.

C documentation : [GLFW_ARROW_CURSOR]

[GLFW_ARROW_CURSOR]: https://www.glfw.org/docs/latest/group__shapes.html#ga8ab0e717245b85506cb0eaefdea39d0a
*/
const ARROW_CURSOR = 0x00036001

/*
The text input I-beam cursor shape.

C documentation : [GLFW_IBEAM_CURSOR]

[GLFW_IBEAM_CURSOR]: https://www.glfw.org/docs/latest/group__shapes.html#ga36185f4375eaada1b04e431244774c86
*/
const IBEAM_CURSOR = 0x00036002

/*
The crosshair cursor shape.

C documentation : [GLFW_CROSSHAIR_CURSOR]

[GLFW_CROSSHAIR_CURSOR]: https://www.glfw.org/docs/latest/group__shapes.html#ga8af88c0ea05ab9e8f9ac1530e8873c22
*/
const CROSSHAIR_CURSOR = 0x00036003

/*
The pointing hand cursor shape.

C documentation : [GLFW_POINTING_HAND_CURSOR]

[GLFW_POINTING_HAND_CURSOR]: https://www.glfw.org/docs/latest/group__shapes.html#gaad01a50929fb515bf27e4462c51f6ed0
*/
const POINTING_HAND_CURSOR = 0x00036004

/*
The horizontal resize/move arrow shape.  This is usually a horizontal
double-headed arrow.

C documentation : [GLFW_RESIZE_EW_CURSOR]

[GLFW_RESIZE_EW_CURSOR]: https://www.glfw.org/docs/latest/group__shapes.html#ga2010a43dc1050a7c9154148a63cf01ad
*/
const RESIZE_EW_CURSOR = 0x00036005

/*
The vertical resize/move shape.  This is usually a vertical double-headed
arrow.

C documentation : [GLFW_RESIZE_NS_CURSOR]

[GLFW_RESIZE_NS_CURSOR]: https://www.glfw.org/docs/latest/group__shapes.html#gaa59214e8cdc8c8adf08fdf125ed68388
*/
const RESIZE_NS_CURSOR = 0x00036006

/*
The top-left to bottom-right diagonal resize/move shape.  This is usually
a diagonal double-headed arrow.

C documentation : [GLFW_RESIZE_NWSE_CURSOR]

# macos

This shape is provided by a private system API and may fail
with [CURSOR_UNAVAILABLE] in the future.

# wayland

This shape is provided by a newer standard not supported by
all cursor themes.

# x11

This shape is provided by a newer standard not supported by all
cursor themes.

[GLFW_RESIZE_NWSE_CURSOR]: https://www.glfw.org/docs/latest/group__shapes.html#gadf2c0a495ec9cef4e1a364cc99aa78da
*/
const RESIZE_NWSE_CURSOR = 0x00036007

/*
The top-right to bottom-left diagonal resize/move shape.  This is usually
a diagonal double-headed arrow.

C documentation : [GLFW_RESIZE_NESW_CURSOR]

# macos

This shape is provided by a private system API and may fail
with [CURSOR_UNAVAILABLE] in the future.

# wayland

This shape is provided by a newer standard not supported by
all cursor themes.

# x11

This shape is provided by a newer standard not supported by all
cursor themes.

[GLFW_RESIZE_NESW_CURSOR]: https://www.glfw.org/docs/latest/group__shapes.html#gab06bba3b407f92807ba9b48de667a323
*/
const RESIZE_NESW_CURSOR = 0x00036008

/*
The omni-directional resize cursor/move shape.  This is usually either
a combined horizontal and vertical double-headed arrow or a grabbing hand.

C documentation : [GLFW_RESIZE_ALL_CURSOR]

[GLFW_RESIZE_ALL_CURSOR]: https://www.glfw.org/docs/latest/group__shapes.html#ga3a5f4811155f95ccafbbb4c9a899fc1d
*/
const RESIZE_ALL_CURSOR = 0x00036009

/*
The operation-not-allowed shape.  This is usually a circle with a diagonal
line through it.

C documentation : [GLFW_NOT_ALLOWED_CURSOR]

# wayland

This shape is provided by a newer standard not supported by
all cursor themes.

# x11

This shape is provided by a newer standard not supported by all
cursor themes.

[GLFW_NOT_ALLOWED_CURSOR]: https://www.glfw.org/docs/latest/group__shapes.html#ga297c503095b034bc8891393b637844b1
*/
const NOT_ALLOWED_CURSOR = 0x0003600A

/*
This is an alias for compatibility with earlier versions.

C documentation : [GLFW_HRESIZE_CURSOR]

[GLFW_HRESIZE_CURSOR]: https://www.glfw.org/docs/latest/group__shapes.html#gabb3eb0109f11bb808fc34659177ca962
*/
const HRESIZE_CURSOR = RESIZE_EW_CURSOR

/*
This is an alias for compatibility with earlier versions.

C documentation : [GLFW_VRESIZE_CURSOR]

[GLFW_VRESIZE_CURSOR]: https://www.glfw.org/docs/latest/group__shapes.html#gaf024f0e1ff8366fb2b5c260509a1fce5
*/
const VRESIZE_CURSOR = RESIZE_NS_CURSOR

/*
This is an alias for compatibility with earlier versions.

C documentation : [GLFW_HAND_CURSOR]

[GLFW_HAND_CURSOR]: https://www.glfw.org/docs/latest/group__shapes.html#ga1db35e20849e0837c82e3dc1fd797263
*/
const HAND_CURSOR = POINTING_HAND_CURSOR

/*
C documentation : [GLFW_CONNECTED]

[GLFW_CONNECTED]: https://www.glfw.org/docs/latest/glfw3_8h.html#abe11513fd1ffbee5bb9b173f06028b9e
*/
const CONNECTED = 0x00040001

/*
C documentation : [GLFW_DISCONNECTED]

[GLFW_DISCONNECTED]: https://www.glfw.org/docs/latest/glfw3_8h.html#aab64b25921ef21d89252d6f0a71bfc32
*/
const DISCONNECTED = 0x00040002

/*
Joystick hat buttons [init hint].

C documentation : [GLFW_JOYSTICK_HAT_BUTTONS]

[init hint]: https://www.glfw.org/docs/latest/group__init.html#gab9c0534709fda03ec8959201da3a9a18
[GLFW_JOYSTICK_HAT_BUTTONS]: https://www.glfw.org/docs/latest/group__init.html#gab9c0534709fda03ec8959201da3a9a18
*/
const JOYSTICK_HAT_BUTTONS = 0x00050001

/*
ANGLE rendering backend [init hint].

C documentation : [GLFW_ANGLE_PLATFORM_TYPE]

[init hint]: https://www.glfw.org/docs/latest/intro_guide.html#GLFW_ANGLE_PLATFORM_TYPE_hint
[GLFW_ANGLE_PLATFORM_TYPE]: https://www.glfw.org/docs/latest/group__init.html#gaec269b24cf549ab46292c0125d8bbdce
*/
const ANGLE_PLATFORM_TYPE = 0x00050002

/*
Platform selection [init hint].

C documentation : [GLFW_PLATFORM]

[init hint]: https://www.glfw.org/docs/latest/group__init.html#ga9d38bf1fdf4f91d6565401734a7cd967
[GLFW_PLATFORM]: https://www.glfw.org/docs/latest/group__init.html#ga9d38bf1fdf4f91d6565401734a7cd967
*/
const PLATFORM = 0x00050003

/*
macOS specific [init hint].

C documentation : [GLFW_COCOA_CHDIR_RESOURCES]

[init hint]: https://www.glfw.org/docs/latest/intro_guide.html#GLFW_COCOA_CHDIR_RESOURCES_hint
[GLFW_COCOA_CHDIR_RESOURCES]: https://www.glfw.org/docs/latest/group__init.html#gab937983147a3158d45f88fad7129d9f2
*/
const COCOA_CHDIR_RESOURCES = 0x00051001

/*
macOS specific [init hint].

C documentation : [GLFW_COCOA_MENUBAR]

[init hint]: https://www.glfw.org/docs/latest/intro_guide.html#GLFW_COCOA_MENUBAR_hint
[GLFW_COCOA_MENUBAR]: https://www.glfw.org/docs/latest/group__init.html#ga71e0b4ce2f2696a84a9b8c5e12dc70cf
*/
const COCOA_MENUBAR = 0x00051002

/*
X11 specific [init hint].

C documentation : [GLFW_X11_XCB_VULKAN_SURFACE]

[init hint]: https://www.glfw.org/docs/latest/intro_guide.html#GLFW_X11_XCB_VULKAN_SURFACE_hint
[GLFW_X11_XCB_VULKAN_SURFACE]: https://www.glfw.org/docs/latest/group__init.html#gaa341e303ebeb8e4199b8ab8be84351f6
*/
const X11_XCB_VULKAN_SURFACE = 0x00052001

/*
Wayland specific [init hint].

C documentation : [GLFW_WAYLAND_LIBDECOR]

[init hint]: https://www.glfw.org/docs/latest/intro_guide.html#GLFW_WAYLAND_LIBDECOR_hint
[GLFW_WAYLAND_LIBDECOR]: https://www.glfw.org/docs/latest/group__init.html#ga2a3f2fd7695902c498b050215b3db452
*/
const WAYLAND_LIBDECOR = 0x00053001

/*
Hint value for [PLATFORM] that enables automatic platform selection.

C documentation : [GLFW_ANY_PLATFORM]

[GLFW_ANY_PLATFORM]: https://www.glfw.org/docs/latest/group__init.html#ga18b2d37374d0dea28cd69194fa85b859
*/
const ANY_PLATFORM = 0x00060000

/*
C documentation : [GLFW_PLATFORM_WIN32]

[GLFW_PLATFORM_WIN32]: https://www.glfw.org/docs/latest/group__init.html#ga8d3d17df2ab57492cef665da52c603a1
*/
const PLATFORM_WIN32 = 0x00060001

/*
C documentation : [GLFW_PLATFORM_COCOA]

[GLFW_PLATFORM_COCOA]: https://www.glfw.org/docs/latest/group__init.html#ga83b18714254f75bc2f0cdbafa0f10b6b
*/
const PLATFORM_COCOA = 0x00060002

/*
C documentation : [GLFW_PLATFORM_WAYLAND]

[GLFW_PLATFORM_WAYLAND]: https://www.glfw.org/docs/latest/group__init.html#gac4b08906a3cbf26c518a4a543eedd740
*/
const PLATFORM_WAYLAND = 0x00060003

/*
C documentation : [GLFW_PLATFORM_X11]

[GLFW_PLATFORM_X11]: https://www.glfw.org/docs/latest/group__init.html#gaf5333f3933e9c248a00cfda6523f386b
*/
const PLATFORM_X11 = 0x00060004

/*
C documentation : [GLFW_PLATFORM_NULL]

[GLFW_PLATFORM_NULL]: https://www.glfw.org/docs/latest/group__init.html#gac06fad5a4866ae7a1d7b2675fac72d7f
*/
const PLATFORM_NULL = 0x00060005

/*
C documentation : [GLFW_DONT_CARE]

[GLFW_DONT_CARE]: https://www.glfw.org/docs/latest/glfw3_8h.html#a7a2edf2c18446833d27d07f1b7f3d571
*/
const DONT_CARE = -1
