// > javac Object.javac
// > java Main.java
// 程序输出效果:
// Person{name='Jack', privName='Mark'}
// Person{name='Kail', privName='null'}
// Person{name='null', privName='null'}
// Jack
// Jack
// Mark
// 21
// 22
// 94759122314
class Main {
    public static void main(String[] args) {
        // 构造Person
		// 全参构造
		Person person1 = new Person("Jack", "Mark");
		System.out.println(person1.toString());

		// 部分参数构造
		Person person2 = new Person("Kail");
		System.out.println(person2.toString());

		// 无参构造
		Person person3 = new Person();
		System.out.println(person3);

		// 访问成员变量, 直接访问或者get方法都可以
		System.out.println(person1.name);
		System.out.println(person1.getName());

		// 私有成员变量无法直接访问, 需要使用get方法
		System.out.println(person1.getPrivName());

		// 访问全局变量(静态成员变量), 操作全局变量
		System.out.println(Person.age);
		Person.age = "22";
		System.out.println(Person.age);

		// 访问常量、常量不可修改
		System.out.println(Person.ID);
    }
}

class Person {
	// 以下部分为定义变量---------------------------------------------
    // 成员变量
    String name;

    // 私有成员变量
    private String privName;

	// 静态成员变量，一般成为全局变量, 可修改
	static String age = "21";

	// 终态成员变量, 一般称之为常量，不可修改，一般作为枚举使用
	// 常量的标准格式为 全大写
	public final static String ID = "94759122314";


	// 以下部分为定义以下部分为定义构造方法
    // 构造方法, 格式: 类名(参数) { ... }
    // 全参构造
    Person(String name, String privName) {
        this.name = name;
        this.privName = privName;
    }

    // 部分参数构造方法
    Person(String name) {
        this.name = name;
    }

    // 无参构造, 默认不需要写, 无参构造默认生成
    Person() { }


	// 以下部分为重写toString方法, 此方法JKD自带存在, 单独使用需要重写
	// @override: 此注解意为重写方法
    @Override
    public String toString() {
        return "Person{" +
                "name='" + name + '\'' +
                ", privName='" + privName + '\'' +
                '}';
    }


    // 以下为get/set方法, 非私有成员变量无需get方法, 因为可以直接访问。
    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }

    public String getPrivName() {
        return privName;
    }

    public void setPrivName(String privName) {
        this.privName = privName;
    }
}
