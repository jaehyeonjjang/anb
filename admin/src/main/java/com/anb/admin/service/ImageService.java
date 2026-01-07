package com.anb.admin.service;

import java.util.Optional;
import java.util.List;
import java.io.File;
import java.nio.file.Files;

import lombok.extern.slf4j.Slf4j;

import lombok.extern.slf4j.Slf4j;

import org.springframework.stereotype.Service;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.transaction.annotation.Transactional;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.data.domain.Pageable;
import org.springframework.data.domain.Page;

import com.anb.admin.domain.Image;
import com.anb.admin.domain.ImageRepository;
import com.anb.admin.util.ImageUtil;

@Slf4j
@Service
public class ImageService {
    @Value("${path.root-path}")
    private String rootPath;

    @Value("${path.upload-path}")
    private String uploadPath;

    @Autowired
    ImageRepository repository;

    private void makeImage(int type, String filename) {
        if (filename == null) {
            return;
        }

        if (filename.equals("")) {
            return;
        }

        String fullFilename = rootPath + "/" + filename.replace("webdata/", "webdata/original/");
        String targetFilename = rootPath + "/" + filename;

        if (type == 3 || type == 4) {
            //ImageUtil.resize(fullFilename, targetFilename, 1600, 1193);
            ImageUtil.resize(fullFilename, targetFilename, 1600, 916, 400);
        } else {
            try {
                Files.copy(new File(fullFilename).toPath(), new File(targetFilename).toPath());
            } catch (Exception e) {
                log.info("File Copy Error : " + fullFilename);
            }
        }
    }

    @Transactional
    public Image insert(Image item) {
        makeImage(item.getType(), item.getFilename());

        return repository.save(item);
    }

    @Transactional
    public Image update(Image item) {
        Optional<Image> opt = repository.findById(item.getId());

        if (!opt.isPresent()) {
            return null;
        }

        makeImage(item.getType(), item.getFilename());

        return repository.save(item);
    }

    @Transactional
    public void delete(Image item) {
        repository.delete(item);
    }

    public Optional<Image> findById(Long id) {
        return repository.findById(id);
    }

    public List<Image> findByApt(Long apt) {
        return repository.findByAptOrderByOrderAscIdAsc(apt);
    }
}
