// Lamdba 表达式 可用于更简便的实现接口，但只限于函数式接口
// 注: 函数式接口代表，此接口除去default默认实现方法与系统默认重写方法（类似tostring），只实现单个接口
// 语法: 
// (参数) -> {
// 		方法体
// };
public class Program {

	public static void main(String[] args) {
		// 无参数、无返回的函数式接口
		InterfaceImpl();

		// Lamdba实现
		Lamdba();
	}

	private static void LamdbaUse() {
		NoneReturnNoneParameter lamdba1 = () -> {
			System.out.println("hello, lamdba");
		};
		lamdba1.use();
	}

	private static void InterfaceImpl() {
		// 使用显示的实现类对象
		SingleReturnSingleParameter parameter1 = new Impl();
		// 匿名内部类，隐式实现
		SingleReturnSingleParameter parameter2 = new SingleReturnSingleParameter() {
			@override
			public int test(int a) {
				return a * a;
			}
		}
		// Lamdba表达式
		SingleReturnSingleParameter parameter3 = a -> a * a;

		// 使用
		System.out.println(parameter1.test(10));
		System.out.println(parameter2.test(10));
		System.out.println(parameter3.test(10));
	}

}


public static void Impl implements SingleReturnSingleParameter {

	@override
	public int test(int a) {
		return a * a;
	}
}


// 此注解可用于声明此接口为函数式接口，与 @override类似
@FunctionInterface
public interface SingleReturnSingleParameter {

	int test();
}

// ---------------------------------- 以下为多个Lamdba表达式函数式接口原型，可直接使用 --------------------------------------
// 无参无返回
@FunctionInterface
public interface NoneReturnNoneParameter {

	void use();
}
