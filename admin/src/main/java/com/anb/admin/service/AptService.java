package com.anb.admin.service;

import java.util.Optional;
import java.util.List;
import java.util.Map;

import org.springframework.stereotype.Service;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.transaction.annotation.Transactional;
import org.springframework.data.domain.Pageable;
import org.springframework.data.domain.Page;
import org.springframework.data.domain.Sort;
import org.springframework.data.domain.PageRequest;
import org.apache.commons.lang3.StringUtils;

import com.anb.admin.domain.Apt;
import com.anb.admin.domain.AptRepository;
import com.anb.admin.domain.Aptgroup;
import com.anb.admin.domain.AptgroupRepository;
import com.anb.admin.domain.AptSpecs;
import com.anb.admin.domain.AptSpecs.SearchKey;
import com.anb.admin.domain.Aptuser;
import com.anb.admin.domain.AptuserRepository;
import com.anb.admin.domain.Aptsubmaster;
import com.anb.admin.domain.AptsubmasterRepository;
import com.anb.admin.domain.Image;
import com.anb.admin.domain.ImageRepository;
import com.anb.admin.domain.Imagefloor;
import com.anb.admin.domain.ImagefloorRepository;
import com.anb.admin.domain.Data;
import com.anb.admin.domain.DataRepository;

@Service
public class AptService {

    @Autowired
    AptRepository repository;

    @Autowired
    AptgroupService aptgroupService;

    @Autowired
    AptuserService aptuserService;

    @Autowired
    AptsubmasterService aptsubmasterService;

    @Autowired
    ImageService imageService;

    @Autowired
    ImagefloorService imagefloorService;

    @Autowired
    DataService dataService;

    public String  getSearch(Apt item) {
        Optional<Aptgroup> opt = aptgroupService.findById(item.getAptgroup());

        if (!opt.isPresent()) {
            return "";
        }

        Aptgroup aptgroup = opt.get();

        String search = aptgroup.getName() + " " + item.getName();

        return search;
    }

    @Transactional
    public Apt insert(Apt item) {
        item.setSearch(getSearch(item));
        return repository.save(item);
    }

    @Transactional
    public Apt update(Apt item) {
        item.setSearch(getSearch(item));
        return repository.save(item);
    }

    @Transactional
    public void delete(Apt item) {
        repository.delete(item);
    }

    public Optional<Apt> findById(Long id) {
        return repository.findById(id);
    }

    public Page<Apt> findAll(Map<SearchKey, Object> searchKeys, String order, int page, int size) {
        Sort sort = null;
        boolean desc = false;
        
        if (StringUtils.isEmpty(order)) {
            order = "id";            
        } else {
            if (StringUtils.right(order, 4).equals("Desc")) {
                order = StringUtils.left(order, order.length() - 4);
                desc = true;
            }
        }

        sort = Sort.by(order);

        if (desc) {
            sort = sort.descending();
        }
        
        Pageable pageable = PageRequest.of(page, size, sort);        

        return searchKeys.isEmpty()
            ? repository.findAll(pageable)
            : repository.findAll(AptSpecs.searchWith(searchKeys), pageable);
    }

    public List<Apt> findByCompany(Long company) {
        return repository.findByCompany(company);
    }

    public List<Apt> findByCompanyAndReport(Long company, int report) {
        return repository.findByCompanyAndReport(company, report);
    }

    public List<Apt> findByCompanyAndStatus(Long company, int status) {
        return repository.findByCompanyAndStatus(company, status);
    }

    public List<Apt> findByAptgroup(Long aptgroup) {
        return repository.findByAptgroup(aptgroup);
    }

    @Transactional
    public Apt copy(Long id) {
        Optional<Apt> opt = findById(id);

        if (!opt.isPresent()) {
            return null;
        }

        Apt old = opt.get();

        Apt item = new Apt();
        item.setAptgroup(old.getAptgroup());
        item.setName(old.getName());
        item.setStartdate(old.getStartdate());
        item.setEnddate(old.getEnddate());
        item.setType(old.getType());
        item.setMaster(old.getMaster());
        item.setStatus(old.getStatus());
        item.setCompany(old.getCompany());

        item.setReport(old.getReport());
        item.setReport1(old.getReport1());
        item.setReport2(old.getReport2());
        item.setReport3(old.getReport3());
        item.setReport4(old.getReport4());
        item.setReport5(old.getReport5());
        item.setReport6(old.getReport6());

        item.setUpdateuser(old.getUpdateuser());
        item.setUser(old.getUser());

        Apt newApt = insert(item);

        List<Aptuser> aptusers = aptuserService.findByApt(id);

        for (Aptuser aptuserItem : aptusers) {
            Aptuser aptuser = new Aptuser();
            aptuser.setApt(newApt.getId());
            aptuser.setUser(aptuserItem.getUser());
            aptuser.setLevel(aptuserItem.getLevel());
            aptuser.setCompany(aptuserItem.getCompany());

            aptuserService.insert(aptuser);
        }

        List<Aptsubmaster> aptsubmasters = aptsubmasterService.findByApt(id);

        for (Aptsubmaster aptsubmasterItem : aptsubmasters) {
            Aptsubmaster aptsubmaster = new Aptsubmaster();
            aptsubmaster.setApt(newApt.getId());
            aptsubmaster.setUser(aptsubmasterItem.getUser());
            aptsubmaster.setLevel(aptsubmasterItem.getLevel());
            aptsubmaster.setCompany(aptsubmasterItem.getCompany());

            aptsubmasterService.insert(aptsubmaster);
        }

        List<Image> images = imageService.findByApt(id);

        for (Image imageItem : images) {
            Image image = new Image();
            image.setApt(newApt.getId());
            image.setName(imageItem.getName());
            image.setLevel(imageItem.getLevel());
            image.setParent(imageItem.getParent());
            image.setLast(imageItem.getLast());
            image.setTitle(imageItem.getTitle());
            image.setType(imageItem.getType());
            image.setFilename(imageItem.getFilename());
            image.setOrder(imageItem.getOrder());

            Image newImage = imageService.insert(image);

            List<Imagefloor> imagefloors = imagefloorService.findByImage(imageItem.getId());
            for (Imagefloor imagefloorItem : imagefloors) {
                Imagefloor imagefloor = new Imagefloor();
                imagefloor.setImage(newImage.getId());
                imagefloor.setName(imagefloorItem.getName());
                imagefloor.setImagename(imagefloorItem.getImagename());
                imagefloor.setTarget(imagefloorItem.getTarget());

                imagefloorService.insert(imagefloor);
            }

            List<Data> datas = dataService.findByImage(imageItem.getId());

            for (Data dataItem : datas) {
                Data data = new Data();
                data.setApt(newApt.getId());      
                data.setImage(newImage.getId());    
                data.setImagetype(dataItem.getImagetype());
                data.setUser(dataItem.getUser());     
                data.setType(dataItem.getType());     
                data.setX(dataItem.getX());        
                data.setY(dataItem.getY());        
                data.setPoint(dataItem.getPoint());    
                data.setNumber(dataItem.getNumber());   
                data.setGroup(dataItem.getGroup());    
                data.setName(dataItem.getName());     
                data.setFault(dataItem.getFault());    
                data.setContent(dataItem.getContent());  
                data.setWidth(dataItem.getWidth());    
                data.setLength(dataItem.getLength());   
                data.setCount(dataItem.getCount());    
                data.setProgress(dataItem.getProgress()); 
                data.setRemark(dataItem.getRemark());   
                data.setImagename(dataItem.getImagename());
                data.setFilename(dataItem.getFilename());
                data.setMemo(dataItem.getMemo());    
                data.setReport(dataItem.getReport());  
                data.setUsermemo(dataItem.getUsermemo());
                data.setAptmemo(dataItem.getAptmemo());

                dataService.insert(data);
            }
        }

        

        return newApt;
    }
}
