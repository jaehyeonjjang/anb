package com.anb.admin.domain;

import java.util.List;
import java.util.ArrayList;
import java.util.Map;

import org.springframework.data.jpa.domain.Specification;

import javax.persistence.criteria.CriteriaBuilder;
import javax.persistence.criteria.CriteriaQuery;
import javax.persistence.criteria.Predicate;
import javax.persistence.criteria.Root;

public class AptSpecs {
    public enum SearchKey {
        STATUS("status"),
        APTGROUP("aptgroup"),
        COMPANY("company"),
        NAME("name"),
        SEARCH("search");

        private final String value;

        SearchKey(String value) {
            this.value = value;
        }

        public String getValue() {
            return value;
        }
    }

    public static Specification<Apt> searchWith(Map<SearchKey, Object> searchKeyword) {
        return (Specification<Apt>) ((root, query, builder) -> {
                List<Predicate> predicate = getPredicateWithKeyword(searchKeyword, root, builder);
                return builder.and(predicate.toArray(new Predicate[0]));
            });
    }

    private static List<Predicate> getPredicateWithKeyword(Map<SearchKey, Object> searchKeyword, Root<Apt> root, CriteriaBuilder builder) {
        List<Predicate> predicate = new ArrayList<>();
        for (SearchKey key : searchKeyword.keySet()) {
            switch (key) {
            case STATUS:
                predicate.add(builder.equal(root.get(key.value), Integer.valueOf(searchKeyword.get(key).toString())));
                break;
            case APTGROUP:
            case COMPANY:
                predicate.add(builder.equal(root.get(key.value), Long.valueOf(searchKeyword.get(key).toString())));
                break;
            case NAME:
                predicate.add(builder.like(root.get(key.value), "%" + searchKeyword.get(key) + "%"));
                break;
            case SEARCH:
                predicate.add(builder.like(root.get(key.value), "%" + searchKeyword.get(key) + "%"));
                break;
            }
        }
        return predicate;
    }
}
