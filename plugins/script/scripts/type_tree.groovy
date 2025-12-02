import java.util.*;

flexibleSearch = spring.getBean("flexibleSearchService");
result = flexibleSearch.search ("select {pk} from {ComposedType}").getResult();

Tree tree = new Tree();
result.each {
    Node node = new Node(it.getCode(), it.getSuperType()?.getCode());
    type = it.getClass().getSimpleName();
    type = type.replace("ComposedTypeModel", "<Composed>");
    type = type.replace("RelationMetaTypeModel", "<Relation>");
    type = type.replace("EnumerationMetaTypeModel", "<ENUM>");
    type = type.replace("TypeModel", "");
    node.setDetails(type);
    tree.getElements().add(node);
}

for (element in tree.getElements()) {
    node1 = tree.find(element.getValue());
    node2 = tree.find(element.getParentValue());
    if (node1 != null) { node1.setParent(node2); }
    if (node2 != null) { node2.addChild(node1); }
    if (element.getParentValue() == null) { root = node1; }
}

int level = 0;
printANode(0, root);
displaySubTree(tree, level, root);

void printANode(level, Node item) {
    print "."*level;
    println item.getValue() + "(" + item.getDetails() + ")";
}

void displaySubTree(Tree tree, int level, Node node)
{
    List<Node> subItems = node.getChildren();
    for (item in subItems) {
        printANode(level+1, item);
        displaySubTree(tree, level+1, item);
    }
}

public class Tree
{
    List<Node> elements;
    public List<Node> getElements() { return elements; }
    public Tree() {
        elements = new ArrayList();
    }
    public void add (Node element) {
        elements.add(element);
    }

    public Node find(String value) {
        for (it in elements) { if (it.getValue() == value) { return it; } }
    }
}

public class Node
{
    private Node parent = null;
    private List<Node> children = null;
    private String details = "";
    private String value = "";
    private String parentValue = "";

    public Node(String value, String parent)
    {
        this.children = new ArrayList<>();
        this.value = value;
        this.parentValue = parent;
    }

    public setDetails(String nodeDetails)
    {
        details = nodeDetails;
    }
    public getDetails() { return details; }

    public List<Node> getChildren() {
        return children;
    }

    public void addChild(Node child)
    {
        children.add(child);
        child.addParent(this);
    }
    public addParent (Node parentNode)
    {

        parent = parentNode;
    }
    public getValue () {
        return value;
    }
    public String getParentValue() {
        return parentValue;
    }
    public setParent(Node node)
    {
        parent = node;
    }
}
