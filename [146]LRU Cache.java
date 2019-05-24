import java.util.HashMap;

//
//Design and implement a data structure for Least Recently Used (LRU) cache. It should support the following operations: get and put.
// 
//
// 
//get(key) - Get the value (will always be positive) of the key if the key exists in the cache, otherwise return -1. 
//put(key, value) - Set or insert the value if the key is not already present. When the cache reached its capacity, it should invalidate the least recently used item before inserting a new item.
// 
//
// Follow up: 
//Could you do both operations in O(1) time complexity? 
//
// Example:
// 
//LRUCache cache = new LRUCache( 2 /* capacity */ );
//
//cache.put(1, 1);
//cache.put(2, 2);
//cache.get(1);       // returns 1
//cache.put(3, 3);    // evicts key 2
//cache.get(2);       // returns -1 (not found)
//cache.put(4, 4);    // evicts key 1
//cache.get(1);       // returns -1 (not found)
//cache.get(3);       // returns 3
//cache.get(4);       // returns 4
// 
//

class LRUCache {

    class Node {
        public Node next;
        public Node prev;
        public int key;
        public int value;
    }
    int cap;

    private HashMap<Integer, Node> nodes = new HashMap<>();
    Node head;
    Node tail;

    public LRUCache(int capacity) {
        cap = capacity;
        int i = 0;
        Node[] nodes = new Node[capacity];
        for (i = 0; i < capacity; i++) {
            nodes[i] = new Node();
        }
        for (i = 0; i < capacity; i++) {
            nodes[i].key = -i;
            nodes[i].value = -1;
            nodes[i].next = nodes[(i + 1) % capacity];
            nodes[i].prev = nodes[i - 1 < 0 ? capacity - 1 : i - 1];
            this.nodes.put(-i, nodes[i]);
        }
        head = nodes[0];
        tail = nodes[capacity - 1];
    }

    public int get(int key) {
        Node node = this.getNode(key);
        return node == null ? -1 : node.value;
    }

    public Node getNode(int key) {
        Node node = nodes.get(key);
        if (node == null) {
            return null;
        } else if (node == head) {
            return node;
        } else {
            node.prev.next = node.next;
            node.next.prev = node.prev;
        }
        node.next = head;
        head.prev = node;
        head = node;
        if (tail == node) {
            tail = node.prev;
        }
        node.prev = tail;
            tail.next = node;

        return node;
    }

    public void put(int key, int value) {
        Node node = this.getNode(key);

        if (node == null) {
            node = this.getNode(tail.key);
        }

        if (node == null) {
            node = tail;
        }
        if (nodes.size() == cap) {
            nodes.remove(node.key);
        }

        node.key = key;
        node.value = value;

        nodes.put(key, node);


    }
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * LRUCache obj = new LRUCache(capacity);
 * int param_1 = obj.get(key);
 * obj.put(key,value);
 */