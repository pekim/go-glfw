# Doxygen tag file

The `tagfile.xml` file in this directory is a Doxygen tag file,
produced by building glfw.
The following steps outline roughly how to reproduce it.

## Clone repo

Clone the glfw git repo, and change in to its directory.

```sh
git clone https://github.com/glfw/glfw
cd glfw
```

## Apply patch

Applying this patch will result in the build producing a tag file.
Given that it's only a one line change it's probably easiest to just
edit the file manually.

```
diff --git a/docs/Doxyfile.in b/docs/Doxyfile.in
index 27c4b3f6..504e04b4 100644
--- a/docs/Doxyfile.in
+++ b/docs/Doxyfile.in
@@ -2357,7 +2357,7 @@ TAGFILES               =
 # tag file that is based on the input files it reads. See section "Linking to
 # external documentation" for more information about the usage of tag files.

-GENERATE_TAGFILE       =
+GENERATE_TAGFILE       = tagfile.xml

 # If the ALLEXTERNALS tag is set to YES, all external class will be listed in
 # the class index. If set to NO, only the inherited external classes will be
```

## Build

```sh
cmake -S . -B build
cd build
make
```

## Copy the tag file

```sh
copy docs/tagfile.xml $GL_PUREGO_REPO/internal/tagfile/tagfile.go
```
