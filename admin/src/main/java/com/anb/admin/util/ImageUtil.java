package com.anb.admin.util;

import java.io.File;
import java.io.InputStream;
import java.io.FileInputStream;

import javax.imageio.ImageIO;
import java.awt.*;
import java.awt.image.BufferedImage;

import lombok.extern.slf4j.Slf4j;

import org.apache.commons.io.FilenameUtils;

@Slf4j
public class ImageUtil {
    public static void resize(String filename, String targetFilename, int width, int height, int extra) {
        File file = new File(filename);

        try {
            InputStream inputStream = new FileInputStream(file);
            BufferedImage inputImage = ImageIO.read(inputStream);

            int imageWidth = inputImage.getWidth();
            int imageHeight = inputImage.getHeight();

            float per = (float)width / (float)height;
            float per2 = (float)imageWidth / (float)imageHeight;

            int targetWidth;
            int targetHeight;

            int x = 0;
            int y = 0;

            if (per > per2) {
                targetHeight = height;
                targetWidth = imageWidth * height / imageHeight;

                x = (width - targetWidth) / 2;
                y = 0;
            } else {
                targetWidth = width;
                targetHeight = imageHeight * width / imageWidth;

                x = 0;
                y = (height - targetHeight) / 2;
            }

            BufferedImage outputImage = new BufferedImage(width + extra, height, inputImage.getType());

            Graphics2D graphics2D = outputImage.createGraphics();
            graphics2D.setColor(Color.WHITE);
            graphics2D.fillRect(0, 0, width + extra, height);
            graphics2D.drawImage(inputImage, x, y, targetWidth, targetHeight, null);
            graphics2D.dispose();

            ImageIO.write(outputImage, FilenameUtils.getExtension(targetFilename), new File(targetFilename));
        } catch (Exception e) {
            log.info("Image reisze error");
        }
    }
}
