go:
	CGO_CPPFLAGS="-I/opt/homebrew/Cellar/opencv/4.10.0_11/include" CGO_LDFLAGS="-L/opt/homebrew/Cellar/opencv/4.10.0_11/lib -lopencv_core -lopencv_face -lopencv_videoio -lopencv_imgproc -lopencv_highgui -lopencv_imgcodecs -lopencv_objdetect -lopencv_features2d -lopencv_video -lopencv_dnn -lopencv_xfeatures2d" go run main.go
py:
	python cut_video_with_python_cv.py
