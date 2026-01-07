package com.anb.admin.domain;

import java.util.List;
import java.util.ArrayList;
import java.util.Map;

import org.springframework.data.jpa.domain.Specification;

import javax.persistence.criteria.CriteriaBuilder;
import javax.persistence.criteria.CriteriaQuery;
import javax.persistence.criteria.Predicate;
import javax.persistence.criteria.Root;

public class ContractSpecs {
    public enum SearchKey {
        STATUS("status"),
        COMPANY("company");

        private final String value;

        SearchKey(String value) {
            this.value = value;
        }

        public String getValue() {
            return value;
        }
    }

    public static Specification<Contract> searchWith(Map<SearchKey, Object> searchKeyword) {
        return (Specification<Contract>) ((root, query, builder) -> {
                List<Predicate> predicate = getPredicateWithKeyword(searchKeyword, root, builder);
                return builder.and(predicate.toArray(new Predicate[0]));
            });
    }

    private static List<Predicate> getPredicateWithKeyword(Map<SearchKey, Object> searchKeyword, Root<Contract> root, CriteriaBuilder builder) {
        List<Predicate> predicate = new ArrayList<>();
        for (SearchKey key : searchKeyword.keySet()) {
            switch (key) {
            case STATUS:
                predicate.add(builder.equal(root.get(key.value), Integer.valueOf(searchKeyword.get(key).toString())));
                break;
            case COMPANY:
                predicate.add(builder.equal(root.get(key.value), Long.valueOf(searchKeyword.get(key).toString())));
                break;
            }
        }
        return predicate;
    }
}
