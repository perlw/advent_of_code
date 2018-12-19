.PHONY: convert
convert:
	ffmpeg -i out.gif -movflags faststart -pix_fmt yuv420p -vf "scale=trunc(iw*8):trunc(ih*8):flags=neighbor,setpts=0.25*PTS" -crf 12 -b:v 500K video.mp4
