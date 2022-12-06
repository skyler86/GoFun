package main

import (
	"github.com/minio/minio-go/v6"
	"log"
)

func main() {
	endpoint := "192.168.56.107:9000"
	accessKeyID := "admin"
	secretAccessKey := "admin123"
	useSSL := false   //注意没有安装证书的填false

	// 初使化minio client对象
	minioClient, err := minio.New(endpoint, accessKeyID, secretAccessKey, useSSL)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("初始化完成！")

	// 链接一个叫 mymusic 的存储桶
	//bucketName := "mymusic"

	//创建一个叫 mymusic 的存储桶
	bucketName := "mymusic"
	location := "cn-north-1"
	err = minioClient.MakeBucket(bucketName, location)
	if err != nil {
		log.Println("创建bucket失败！")
		// 检查存储桶是否已经存在
		exists, err := minioClient.BucketExists(bucketName)
		if err == nil && exists {
			log.Printf("We already own %s\n", bucketName)
		} else {
			log.Println("打印失败！")
			log.Fatalln(err)
		}
	}else {
		log.Printf("Successfully created %s\n", bucketName)

		// 定义上传文件的信息
		objectName := "mp3音乐/执迷不悟.mp3"
		filePath := "C:\\本地存储库\\Music\\华语经典\\柏静-执迷不悟.mp3"
		contentType := "audio/mp3"

		// 使用 FPutObject 给存储桶上传一个文件
		n, err := minioClient.FPutObject(bucketName, objectName, filePath, minio.PutObjectOptions{ContentType: contentType})
		if err != nil {
			log.Fatalln(err)
		}

		log.Printf("Successfully uploaded %s of size %d\n", objectName, n)
	}
}
